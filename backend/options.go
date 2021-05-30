package startpage

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/ecnepsnai/startpage/util"
)

// SPOptions describes options for the startpage server
type SPOptions struct {
	General GeneralOptions
	Modules ModuleOptions
}

// GeneralOptions describes general options
type GeneralOptions struct {
	ServerURL string
}

// ModuleOptions describes module options
type ModuleOptions struct {
	Weather ModuleWeatherOptions
}

// ModuleWeatherOptions describes module options
type ModuleWeatherOptions struct {
	Latitude  float64
	Longitude float64
	APIKey    string
}

// Options the global options
var Options *SPOptions
var optionsLock = sync.Mutex{}

// LoadOptions load SP options
func LoadOptions() {
	defaults := SPOptions{
		General: GeneralOptions{
			ServerURL: fmt.Sprintf("http://%s:%d", util.SystemHostname(), routerPort()),
		},
		Modules: ModuleOptions{
			Weather: ModuleWeatherOptions{
				Latitude:  49.2829766,
				Longitude: -123.1204358,
				APIKey:    os.Getenv("OW_API_KEY"),
			},
		},
	}

	if !FileExists(path.Join(Directories.Data, "furdl.conf")) {
		Options = &defaults
		if err := Options.Save(); err != nil {
			log.Fatal("Error setting default options: %s", err.Error())
		}
	} else {
		f, err := os.OpenFile(path.Join(Directories.Data, "furdl.conf"), os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatal("Error opening config file: %s", err.Error())
		}
		defer f.Close()
		options := defaults
		if err := json.NewDecoder(f).Decode(&options); err != nil {
			log.Fatal("Error decoding options: %s", err.Error())
		}
		Options = &options
	}
}

// Save save the options to disk
func (o *SPOptions) Save() error {
	optionsLock.Lock()
	defer optionsLock.Unlock()

	f, err := os.OpenFile(path.Join(Directories.Data, "furdl.conf"), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Error("Error opening config file: %s", err.Error())
		return err
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(o); err != nil {
		log.Error("Error encoding options: %s", err.Error())
		return err
	}

	Options = o
	return nil
}
