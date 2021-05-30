package weather

import (
	"github.com/briandowns/openweathermap"
)

// Sample describes a sample of weather data
type Sample struct {
	Summary     string
	Description string
	IconCode    string
	Temp        float64
	TempHigh    float64
	TempLow     float64
}

// Forecast describes the forecast
type Forecast struct {
	Samples []Sample
}

func (i *Instance) getCurrent() (*Sample, error) {
	api, err := openweathermap.NewCurrent("C", "EN", i.options.APIKey)
	if err != nil {
		return nil, err
	}

	if err := api.CurrentByCoordinates(&openweathermap.Coordinates{Longitude: i.options.Longitude, Latitude: i.options.Latitude}); err != nil {
		return nil, err
	}

	current := Sample{
		Summary:     api.Weather[0].Main,
		Description: api.Weather[0].Description,
		IconCode:    api.Weather[0].Icon,
		Temp:        api.Main.Temp,
		TempHigh:    api.Main.TempMax,
		TempLow:     api.Main.TempMin,
	}
	log.PInfo("Loaded current weather", map[string]interface{}{
		"summary":     current.Summary,
		"description": current.Description,
		"icon_code":   current.IconCode,
		"temp":        current.Temp,
		"temp_high":   current.TempHigh,
		"temp_low":    current.TempLow,
	})

	return &current, nil
}

func (i *Instance) getForecast() (*Forecast, error) {
	api, err := openweathermap.NewForecast("5", "C", "EN", i.options.APIKey)
	if err != nil {
		return nil, err
	}

	data := openweathermap.Forecast5WeatherData{}
	api.ForecastWeatherJson = &data

	if err := api.DailyByCoordinates(&openweathermap.Coordinates{Longitude: i.options.Longitude, Latitude: i.options.Latitude}, 5); err != nil {
		return nil, err
	}

	samples := make([]Sample, len(data.List))
	for i, day := range data.List {
		samples[i] = Sample{
			Summary:     day.Weather[0].Main,
			Description: day.Weather[0].Description,
			IconCode:    day.Weather[0].Icon,
			Temp:        day.Main.Temp,
			TempHigh:    day.Main.TempMax,
			TempLow:     day.Main.TempMin,
		}
	}
	log.PInfo("Loaded forecast weather", map[string]interface{}{
		"num_days": len(samples),
	})

	return &Forecast{samples}, nil
}
