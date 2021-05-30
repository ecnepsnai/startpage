package startpage

// This file is was generated automatically by Codegen v1.6.0
// Do not make changes to this file as they will be lost

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func getAPIOperatingDir() string {
	ex, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to determine working directory: %s\n"+err.Error())
		os.Exit(1)
	}
	return filepath.Dir(ex)
}

var operatingDirectory = getAPIOperatingDir()
var dataDirectory = getAPIOperatingDir()

type apiDirectories struct {
	Base string

	Data string

	Bookmarks string

	Logs string
}

// Directories absolute paths of API related directires.
var Directories = apiDirectories{}

func fsSetup() {
	Directories = apiDirectories{
		Base: operatingDirectory,

		Data: path.Join(dataDirectory, "data"),

		Bookmarks: path.Join(dataDirectory, "data", "bookmarks"),

		Logs: path.Join(dataDirectory, "logs"),
	}

	MakeDirectoryIfNotExist(Directories.Data)

	MakeDirectoryIfNotExist(Directories.Bookmarks)

	MakeDirectoryIfNotExist(Directories.Logs)

}

// DirectoryExists does the given directory exist (and is it a directory)
func DirectoryExists(directoryPath string) bool {
	stat, err := os.Stat(directoryPath)
	return err == nil && stat.IsDir()
}

// MakeDirectoryIfNotExist make the given directory if it does not exist
func MakeDirectoryIfNotExist(directoryPath string) {
	if !DirectoryExists(directoryPath) {
		os.MkdirAll(directoryPath, 0755)
	}
}

// FileExists does the given file exist
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Fprintf(os.Stderr, "Error stat-ing file '%s': %s", filePath, err.Error())
	return false
}
