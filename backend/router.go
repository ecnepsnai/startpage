package startpage

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ecnepsnai/web"
)

var StaticPath = "./static/build"
var HTTPBindAddress = "localhost:8080"

type apiHandle struct{}
type viewHandle struct{}

func StartRouter() {
	options := web.HandleOptions{}

	api := &apiHandle{}
	view := &viewHandle{}

	server := web.New(HTTPBindAddress)

	server.HTTP.Static("/static/*filepath", StaticPath)
	server.HTTP.Static(fmt.Sprintf("/startpage%s/*filepath", StartpageVersion), StaticPath)
	server.HTTP.GET("/", view.Index, options)

	// Module - PotD
	server.API.GET("/api/modules/potd/info", api.ModulePotdInfo, options)
	server.HTTP.GET("/api/modules/potd/picture", view.ModulePotdPicture, options)
	// Module - MEDailyDeal
	server.API.GET("/api/modules/medailydeal/info", api.ModuleMEDailyDealInfo, options)
	server.HTTP.GET("/api/modules/medailydeal/picture", view.ModuleMEDailyDealPicture, options)
	// Module - Weather
	server.API.GET("/api/modules/weather/info", api.ModuleWeatherInfo, options)

	// Bookmark
	server.API.GET("/api/bookmarks", api.BookmarkGetAll, options)
	server.API.PUT("/api/bookmarks/bookmark", api.BookmarkNew, options)
	server.API.GET("/api/bookmarks/bookmark/:id", api.BookmarkGet, options)
	server.HTTP.GET("/api/bookmarks/bookmark/:id/picture", view.BookmarkGetPicture, options)
	server.API.POST("/api/bookmarks/bookmark/:id", api.BookmarkEdit, options)
	server.API.POST("/api/bookmarks/bookmark/:id/picture", api.BookmarkEditPicture, options)
	server.API.DELETE("/api/bookmarks/bookmark/:id", api.BookmarkDelete, options)

	server.Start()
}
func routerPort() uint16 {
	c := strings.Split(HTTPBindAddress, ":")
	p, err := strconv.ParseUint(c[len(c)-1], 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(p)
}
