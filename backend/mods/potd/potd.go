// Package potd is a startpage module for fetching the photo of the day
package potd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/ecnepsnai/logtic"
	"github.com/ecnepsnai/startpage/mods/potd/bing"
	"github.com/ecnepsnai/startpage/util"
)

var log = logtic.Connect("mod:potd")

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

	if i.getCachedPicture() != nil {
		return nil
	}
	if _, err := i.getNewPicture(); err != nil {
		return err
	}
	return nil
}

// Get will return the cached picture, if there is any
func (i *Instance) Get() *Picture {
	i.lock.Lock()
	defer i.lock.Unlock()

	if cached := i.getCachedPicture(); cached != nil {
		return cached
	}

	if latest, _ := i.getNewPicture(); latest != nil {
		return latest
	}

	return nil
}

func (i *Instance) getCachedPicture() *Picture {
	cacheConfigPath := path.Join(i.dataDir, "potd.json")
	if !util.FileExists(cacheConfigPath) {
		return nil
	}

	picture := Picture{}
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
	if err := json.NewDecoder(f).Decode(&picture); err != nil {
		log.PError("Error decoding cache file", map[string]interface{}{
			"file_path": cacheConfigPath,
			"error":     err.Error(),
		})
		os.Remove(cacheConfigPath)
		return nil
	}

	if time.Since(picture.EndDate) > 0 {
		return nil
	}

	if !util.FileExists(picture.Path) {
		log.PWarn("Cached picture does not exist", map[string]interface{}{
			"picture_path": picture.Path,
		})
		return nil
	}

	log.PDebug("Loaded cached picture", map[string]interface{}{
		"url":        picture.URL,
		"path":       picture.Path,
		"start_date": picture.StartDate,
		"end_date":   picture.EndDate,
	})
	return &picture
}

func (i *Instance) getNewPicture() (*Picture, error) {
	bPicture, err := bing.GetPicture()
	if err != nil {
		log.PError("Error getting picture from bing", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}

	picturePath := path.Join(i.dataDir, fmt.Sprintf("%d.jpg", bPicture.StartDate))

	f, err := os.OpenFile(picturePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.PError("Error opening picture path", map[string]interface{}{
			"picture_path": picturePath,
			"error":        err.Error(),
		})
		return nil, err
	}
	defer f.Close()

	size, err := bPicture.Download(f)
	if err != nil {
		log.PError("Error downloading picture", map[string]interface{}{
			"url":          bPicture.URL(),
			"picture_path": picturePath,
			"error":        err.Error(),
		})
		return nil, err
	}

	start, err := time.Parse("20060102", fmt.Sprintf("%d", bPicture.StartDate))
	if err != nil {
		log.PError("Error parsing picture start date", map[string]interface{}{
			"start_date": bPicture.StartDate,
			"error":      err.Error(),
		})
		return nil, err
	}
	end, err := time.Parse("20060102", fmt.Sprintf("%d", bPicture.EndDate))
	if err != nil {
		log.PError("Error parsing picture end date", map[string]interface{}{
			"end_date": bPicture.EndDate,
			"error":    err.Error(),
		})
		return nil, err
	}

	picture := Picture{
		URL:       bPicture.URL(),
		Path:      picturePath,
		Size:      size,
		StartDate: start,
		EndDate:   end,
	}

	cacheConfigPath := path.Join(i.dataDir, "potd.json")
	cf, err := os.OpenFile(cacheConfigPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.PError("Error opening cache file", map[string]interface{}{
			"file_path": cacheConfigPath,
			"error":     err.Error(),
		})
		return nil, nil
	}
	defer cf.Close()
	if err := json.NewEncoder(cf).Encode(&picture); err != nil {
		log.PError("Error writing cache file", map[string]interface{}{
			"file_path": cacheConfigPath,
			"error":     err.Error(),
		})
		return nil, nil
	}

	log.PInfo("Updated picture", map[string]interface{}{
		"url":        picture.URL,
		"path":       picture.Path,
		"start_date": picture.StartDate,
		"end_date":   picture.EndDate,
	})
	return &picture, nil
}

// Teardown will tear down and close any open files
func (i *Instance) Teardown() {}
