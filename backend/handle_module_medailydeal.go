package startpage

import (
	"fmt"

	"github.com/ecnepsnai/web"
)

func (a *apiHandle) ModuleMEDailyDealInfo(request web.Request) (interface{}, *web.Error) {
	dailyDeal := moduleMEDailyDeal.Get()
	if dailyDeal == nil {
		return nil, web.ValidationError("no cached daily deal")
	}

	return dailyDeal, nil
}

func (v *viewHandle) ModuleMEDailyDealPicture(request web.Request, writer web.Writer) web.Response {
	dailyDeal := moduleMEDailyDeal.Get()
	if dailyDeal == nil {
		return web.Response{
			Status: 404,
		}
	}

	f, err := dailyDeal.Open()
	if err != nil {
		return web.Response{
			Status: 500,
		}
	}

	return web.Response{
		Status: 200,
		Reader: f,
		Headers: map[string]string{
			"Content-Type":   "image/jpeg",
			"Expires":        dailyDeal.Expires.Format("Mon, 2 Jan 2006 15:04:00 MST"),
			"Content-Length": fmt.Sprintf("%d", dailyDeal.ImageSize),
		},
	}
}
