package startpage

import (
	"fmt"

	"github.com/ecnepsnai/web"
)

func (a *apiHandle) ModulePotdInfo(request web.Request) (interface{}, *web.Error) {
	picture := modulePOTD.Get()
	if picture == nil {
		return nil, web.ValidationError("no cached picture")
	}

	return picture, nil
}

func (v *viewHandle) ModulePotdPicture(request web.Request, writer web.Writer) web.Response {
	picture := modulePOTD.Get()
	if picture == nil {
		return web.Response{
			Status: 404,
		}
	}

	f, err := picture.Open()
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
			"Expires":        picture.EndDate.Format("Mon, 2 Jan 2006 15:04:00 MST"),
			"Content-Length": fmt.Sprintf("%d", picture.Size),
		},
	}
}
