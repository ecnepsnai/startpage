// Package bing provides a method to fetch the image of the day from bing.com. Not affiliated with Microsoft.
package bing

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// Picture describes the structure of a picture from bing.com
type Picture struct {
	XMLName       xml.Name `xml:"image"`
	StartDate     int      `xml:"startdate"`
	FullStartDate int      `xml:"fullstartdate"`
	EndDate       int      `xml:"enddate"`
	URLPath       string   `xml:"url"`
}

func (p *Picture) URL() string {
	return "https://www.bing.com" + p.URLPath
}

// Download will download the picture to the provided reader
func (p *Picture) Download(w io.Writer) (uint64, error) {
	req, err := http.NewRequest("GET", p.URL(), nil)
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

type bingImgResp struct {
	XMLName xml.Name  `xml:"images"`
	Images  []Picture `xml:"image"`
}

// GetPicture will fetch the daily picture from bing.com
func GetPicture() (*Picture, error) {
	resp, err := http.Get("https://www.bing.com/hpimagearchive.aspx?format=xml&idx=0&n=1&mbl=1&mkt=en-ww")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	results := bingImgResp{}

	if err := xml.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	if len(results.Images) < 0 {
		return nil, fmt.Errorf("no images")
	}

	image := results.Images[0]
	return &image, nil
}
