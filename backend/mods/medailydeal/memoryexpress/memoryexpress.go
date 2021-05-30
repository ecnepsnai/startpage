package memoryexpress

import (
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

// DailyDeal describes a daily deal
type DailyDeal struct {
	SKU          string
	Title        string
	URL          string
	RegularPrice string
	SalePrice    string
	ImageURL     string
}

var selectorHighlightDeal = cascadia.MustCompile("div.c-shhp-highlight__deal")
var selectorDailyDealName = cascadia.MustCompile("div.c-shhp-daily-deal__details-name")
var selectorDailyDealPrice = cascadia.MustCompile("div.c-shhp-daily-deal__details-price")
var selectorDailyDealImage = cascadia.MustCompile("div.c-shhp-daily-deal__image")
var selectorAnchor = cascadia.MustCompile("a")
var selectorDiv = cascadia.MustCompile("div")
var selectorSpan = cascadia.MustCompile("span")

// GetDailyDeal find the deal of the day from the given reader
func GetDailyDeal(r io.Reader) (*DailyDeal, error) {
	dom, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	dealNode := selectorHighlightDeal.MatchFirst(dom)
	if dealNode == nil {
		return nil, nil
	}

	dealName := selectorDailyDealName.MatchFirst(dealNode)
	if dealName == nil {
		return nil, nil
	}
	nameAnchor := selectorAnchor.MatchFirst(dealName)
	if nameAnchor == nil {
		return nil, nil
	}
	title := cleanString(nameAnchor.FirstChild.Data)
	href := nodeAttrMap(nameAnchor.Attr)["href"]
	SKU := strings.Split(href, "/")[2]

	dealPrice := selectorDailyDealPrice.MatchFirst(dealNode)
	if dealPrice == nil {
		return nil, nil
	}
	regularPrice, salePrice := getPrices(dealPrice)
	if regularPrice == "" {
		return nil, nil
	}

	dealImage := selectorDailyDealImage.MatchFirst(dealNode)
	if dealImage == nil {
		return nil, nil
	}

	imageURL := getPicture(dealImage)

	dailyDeal := DailyDeal{
		SKU:          SKU,
		Title:        title,
		URL:          "https://www.memoryexpress.com" + href,
		RegularPrice: regularPrice,
		SalePrice:    salePrice,
		ImageURL:     imageURL,
	}
	return &dailyDeal, nil
}

// DownloadPicture will download the picture to the provided reader
func (p *DailyDeal) DownloadPicture(w io.Writer) (uint64, error) {
	req, err := http.NewRequest("GET", p.ImageURL, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Accept", "image/jpeg")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	written, err := io.Copy(w, resp.Body)
	if err != nil {
		return 0, err
	}
	return uint64(written), nil
}

func getPicture(node *html.Node) string {
	img := cascadia.MustCompile("img").MatchFirst(node)
	if img == nil {
		return ""
	}
	return nodeAttrMap(img.Attr)["src"]
}

func getPrices(node *html.Node) (string, string) {
	divs := selectorDiv.MatchAll(node)
	if len(divs) != 3 {
		return "", ""
	}
	regularPriceNode := divs[1]
	regularSpans := selectorSpan.MatchAll(regularPriceNode)
	if len(regularSpans) != 2 {
		return "", ""
	}
	regularPrice := cleanString(regularSpans[1].FirstChild.Data)

	salePriceNode := divs[2]
	saleSpans := selectorSpan.MatchAll(salePriceNode)
	if len(saleSpans) != 3 {
		return "", ""
	}
	salePrice := cleanString(saleSpans[1].FirstChild.Data) + cleanString(saleSpans[2].FirstChild.Data)

	return regularPrice, salePrice
}

func cleanString(in string) string {
	out := in
	out = regexp.MustCompile("  +").ReplaceAllString(out, "")
	out = strings.ReplaceAll(out, "\n", "")
	return out
}

func nodeAttrMap(attr []html.Attribute) map[string]string {
	m := map[string]string{}
	for _, attr := range attr {
		m[attr.Key] = attr.Val
	}
	return m
}
