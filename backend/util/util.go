package util

import (
	"os"
)

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
	return false
}

// SystemHostname return the system hostname or panic
func SystemHostname() string {
	h, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return h
}
