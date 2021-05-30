package startpage

import (
	"path"
	"sync"
	"time"

	"github.com/ecnepsnai/startpage/mods/medailydeal"
	"github.com/ecnepsnai/startpage/mods/potd"
	"github.com/ecnepsnai/startpage/mods/weather"
)

var moduleMEDailyDeal *medailydeal.Instance
var modulePOTD *potd.Instance
var moduleWeather *weather.Instance

func ModuleSetup() {
	log.Debug("Starting module setup")

	imedailydeal, err := medailydeal.Setup(path.Join(Directories.Data, "medailydeal"), nil)
	if err != nil {
		log.PFatal("Error initalizing module", map[string]interface{}{
			"module": "potd",
			"error":  err.Error(),
		})
	}
	moduleMEDailyDeal = imedailydeal

	ipotd, err := potd.Setup(path.Join(Directories.Data, "potd"), nil)
	if err != nil {
		log.PFatal("Error initalizing module", map[string]interface{}{
			"module": "potd",
			"error":  err.Error(),
		})
	}
	modulePOTD = ipotd

	iweather, err := weather.Setup(path.Join(Directories.Data, "weather"), &weather.Options{
		Latitude:  Options.Modules.Weather.Latitude,
		Longitude: Options.Modules.Weather.Longitude,
		APIKey:    Options.Modules.Weather.APIKey,
	})
	if err != nil {
		log.PFatal("Error initalizing module", map[string]interface{}{
			"module": "weather",
			"error":  err.Error(),
		})
	}
	moduleWeather = iweather

	log.Debug("Module setup complete")
}

func ModuleRefresh() error {
	start := time.Now()
	log.Debug("Starting module refresh")

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		if err := moduleMEDailyDeal.Refresh(); err != nil {
			log.PError("Error refreshing module", map[string]interface{}{
				"module": "potd",
				"error":  err.Error(),
			})
		}
	}()

	go func() {
		defer wg.Done()
		if err := modulePOTD.Refresh(); err != nil {
			log.PError("Error refreshing module", map[string]interface{}{
				"module": "potd",
				"error":  err.Error(),
			})
		}
	}()

	go func() {
		defer wg.Done()
		if err := moduleWeather.Refresh(); err != nil {
			log.PError("Error refreshing module", map[string]interface{}{
				"module": "potd",
				"error":  err.Error(),
			})
		}
	}()

	wg.Wait()

	log.Debug("Module refresh completed in %s", time.Since(start))
	return nil
}

func ModuleTeardown() {
	moduleMEDailyDeal.Teardown()
	modulePOTD.Teardown()
	moduleWeather.Teardown()
}
