package weather

import (
	"encoding/json"
	"os"
	"path"
	"time"

	"github.com/ecnepsnai/startpage/util"
)

type Weather struct {
	Current  Sample
	Forecast Forecast
	Expires  time.Time
}

func (i *Instance) getCachedWeather() *Weather {
	cachePath := path.Join(i.dataDir, "weather.json")
	if !util.FileExists(cachePath) {
		return nil
	}

	weatherCache := Weather{}
	f, err := os.OpenFile(cachePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.PError("Error reading cache file", map[string]interface{}{
			"file_path": cachePath,
			"error":     err.Error(),
		})
		os.Remove(cachePath)
		return nil
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&weatherCache); err != nil {
		log.PError("Error decoding cache file", map[string]interface{}{
			"file_path": cachePath,
			"error":     err.Error(),
		})
		os.Remove(cachePath)
		return nil
	}

	if time.Since(weatherCache.Expires) > 0 {
		return nil
	}

	log.Debug("Loaded cached weather data")
	return &weatherCache
}

func (i *Instance) getWeather() (*Weather, error) {
	current, err := i.getCurrent()
	if err != nil {
		log.PError("Error getting current weather", map[string]interface{}{
			"error": err.Error(),
		})
	}
	forecast, err := i.getForecast()
	if err != nil {
		log.PError("Error getting forecast", map[string]interface{}{
			"error": err.Error(),
		})
	}
	expires := time.Now().Add(1 * time.Hour)

	weather := Weather{
		Current:  *current,
		Forecast: *forecast,
		Expires:  expires,
	}

	cachePath := path.Join(i.dataDir, "weather.json")
	cf, err := os.OpenFile(cachePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.PError("Error opening cache file", map[string]interface{}{
			"file_path": cachePath,
			"error":     err.Error(),
		})
		return nil, nil
	}
	defer cf.Close()
	if err := json.NewEncoder(cf).Encode(&weather); err != nil {
		log.PError("Error writing cache file", map[string]interface{}{
			"file_path": cachePath,
			"error":     err.Error(),
		})
		return nil, nil
	}

	log.Debug("Cached weather data")
	return &weather, nil
}
