// Package weather is a startpage module for fetching the current weather and 7-day forecast
package weather

import (
	"fmt"
	"sync"

	"github.com/ecnepsnai/logtic"
	"github.com/ecnepsnai/startpage/util"
)

var log = logtic.Connect("mod:weather")

// Options describes options for this module
type Options struct {
	Latitude  float64
	Longitude float64
	APIKey    string
}

// Instance describes an instance of this module
type Instance struct {
	lock    *sync.Mutex
	dataDir string
	options Options
}

// Setup will prepare this module using the provided options
func Setup(dataDir string, options *Options) (*Instance, error) {
	if options == nil {
		return nil, fmt.Errorf("options required")
	}
	if options.APIKey == "" {
		return nil, fmt.Errorf("api key required")
	}
	util.MakeDirectoryIfNotExist(dataDir)
	log.PDebug("Module init", map[string]interface{}{
		"data_dir": dataDir,
		"options":  options,
	})
	return &Instance{
		lock:    &sync.Mutex{},
		dataDir: dataDir,
		options: *options,
	}, nil
}

// Refresh will refresh all data for this module
func (i *Instance) Refresh() error {
	i.lock.Lock()
	defer i.lock.Unlock()

	if i.getCachedWeather() != nil {
		return nil
	}
	if _, err := i.getWeather(); err != nil {
		return err
	}
	return nil
}

// Get will return the cached weather, if there is any
func (i *Instance) Get() *Weather {
	i.lock.Lock()
	defer i.lock.Unlock()

	if cached := i.getCachedWeather(); cached != nil {
		return cached
	}

	if latest, _ := i.getWeather(); latest != nil {
		return latest
	}

	return nil
}

// Teardown will tear down and close any open files
func (i *Instance) Teardown() {}
