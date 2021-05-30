// Package weather is a startpage module for fetching the memory express daily deal
package medailydeal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"github.com/ecnepsnai/logtic"
	"github.com/ecnepsnai/startpage/mods/medailydeal/memoryexpress"
	"github.com/ecnepsnai/startpage/util"
)

var log = logtic.Connect("mod:medailydeal")

// Options describes options for this module
type Options struct{}

// Instance describes an instance of this module
type Instance struct {
	lock    *sync.Mutex
	dataDir string
}

// Setup will prepare this module using the provided options
func Setup(dataDir string, options *Options) (*Instance, error) {
	util.MakeDirectoryIfNotExist(dataDir)
	log.PDebug("Module init", map[string]interface{}{
		"data_dir": dataDir,
		"options":  options,
	})
	i := Instance{
		lock:    &sync.Mutex{},
		dataDir: dataDir,
	}
	return &i, nil
}

// Refresh will refresh all data for this module
func (i *Instance) Refresh() error {
	i.lock.Lock()
	defer i.lock.Unlock()

	if i.getCachedDailyDeal() != nil {
		return nil
	}
	if _, err := i.getNewDailyDeal(); err != nil {
		return err
	}
	return nil
}

// Get will return the cached daily deal, if there is any
func (i *Instance) Get() *DailyDeal {
	i.lock.Lock()
	defer i.lock.Unlock()

	if cached := i.getCachedDailyDeal(); cached != nil {
		return cached
	}

	if latest, _ := i.getNewDailyDeal(); latest != nil {
		return latest
	}

	return nil
}

func (i *Instance) getCachedDailyDeal() *DailyDeal {
	cacheConfigPath := path.Join(i.dataDir, "deal.json")
	if !util.FileExists(cacheConfigPath) {
		return nil
	}

	dailyDeal := DailyDeal{}
	f, err := os.OpenFile(cacheConfigPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.PError("Error reading cache file", map[string]interface{}{
			"file_path": cacheConfigPath,
			"error":     err.Error(),
		})
		os.Remove(cacheConfigPath)
		return nil
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&dailyDeal); err != nil {
		log.PError("Error decoding cache file", map[string]interface{}{
			"file_path": cacheConfigPath,
			"error":     err.Error(),
		})
		os.Remove(cacheConfigPath)
		return nil
	}

	if time.Since(dailyDeal.Expires) > 0 {
		return nil
	}

	log.PDebug("Loaded cached picture", map[string]interface{}{
		"sku":           dailyDeal.SKU,
		"title":         dailyDeal.Title,
		"url":           dailyDeal.URL,
		"regular_price": dailyDeal.RegularPrice,
		"sale_price":    dailyDeal.SalePrice,
		"image_path":    dailyDeal.ImagePath,
		"image_size":    dailyDeal.ImageSize,
		"expires":       dailyDeal.Expires,
	})
	return &dailyDeal
}

func (i *Instance) getNewDailyDeal() (*DailyDeal, error) {
	t, _ := time.Parse("20060102MST", time.Now().Format("20060102MST"))
	expires := t.AddDate(0, 0, 1)

	resp, err := http.Get("https://www.memoryexpress.com/")
	if err != nil {
		log.PError("Error getting memory express index", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	defer resp.Body.Close()
	deal, err := memoryexpress.GetDailyDeal(resp.Body)
	if err != nil {
		log.PError("Error getting daily deal", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	if deal == nil {
		log.PError("No daily deal returned", map[string]interface{}{
			"error": "no data",
		})
		return nil, fmt.Errorf("no data")
	}

	picturePath := path.Join(i.dataDir, fmt.Sprintf("%s.jpg", deal.SKU))
	f, err := os.OpenFile(picturePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.PError("Error opening picture path", map[string]interface{}{
			"picture_path": picturePath,
			"error":        err.Error(),
		})
		return nil, err
	}
	defer f.Close()
	imageSize, err := deal.DownloadPicture(f)
	if err != nil {
		log.PError("Error downloading deal picture", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}

	dailyDeal := DailyDeal{
		SKU:          deal.SKU,
		Title:        deal.Title,
		URL:          deal.URL,
		RegularPrice: deal.RegularPrice,
		SalePrice:    deal.SalePrice,
		ImagePath:    picturePath,
		ImageSize:    imageSize,
		Expires:      expires,
	}

	cacheConfigPath := path.Join(i.dataDir, "deal.json")
	cf, err := os.OpenFile(cacheConfigPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.PError("Error opening cache file", map[string]interface{}{
			"file_path": cacheConfigPath,
			"error":     err.Error(),
		})
		return nil, nil
	}
	defer cf.Close()
	if err := json.NewEncoder(cf).Encode(&dailyDeal); err != nil {
		log.PError("Error writing cache file", map[string]interface{}{
			"file_path": cacheConfigPath,
			"error":     err.Error(),
		})
		return nil, nil
	}

	log.PInfo("Updated picture", map[string]interface{}{
		"sku":           dailyDeal.SKU,
		"title":         dailyDeal.Title,
		"url":           dailyDeal.URL,
		"regular_price": dailyDeal.RegularPrice,
		"sale_price":    dailyDeal.SalePrice,
		"image_path":    dailyDeal.ImagePath,
		"image_size":    dailyDeal.ImageSize,
		"expires":       dailyDeal.Expires,
	})
	return &dailyDeal, nil
}

// Teardown will tear down and close any open files
func (i *Instance) Teardown() {}
