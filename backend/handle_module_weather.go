package startpage

import (
	"github.com/ecnepsnai/web"
)

func (a *apiHandle) ModuleWeatherInfo(request web.Request) (interface{}, *web.Error) {
	weather := moduleWeather.Get()
	if weather == nil {
		return nil, web.ValidationError("no cached weather")
	}

	return weather, nil
}
