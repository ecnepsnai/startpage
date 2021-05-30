package util_test

import (
	"os"
	"path"
	"testing"

	"github.com/ecnepsnai/startpage/util"
)

func TestDirectoryExists(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "startpage")
	if err != nil {
		panic(err)
	}

	if !util.DirectoryExists(dir) {
		t.Errorf("Incorrect result for DirectoryExists (false when expected true)")
	}
	if util.DirectoryExists("/this/path/should/not/exist/on/your/computer") {
		t.Errorf("Incorrect result for DirectoryExists (true when expected false)")
	}
}

func TestMakeDirectoryIfNotExist(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "startpage")
	if err != nil {
		panic(err)
	}
	dirPath := path.Join(dir, "startpage")

	var expectedPermissions os.FileMode = os.ModeDir | (04<<6 | 02<<6) | 01<<6 // dwrx------
	if err := os.Mkdir(dirPath, expectedPermissions); err != nil {
		panic(err)
	}

	util.MakeDirectoryIfNotExist(dirPath)

	info, err := os.Stat(dirPath)
	if err != nil {
		panic(err)
	}
	if info.Mode() != expectedPermissions {
		t.Errorf("Unexpected directory permissions. Expected %v got %v", expectedPermissions, info.Mode())
	}
}

func TestFileExists(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "startpage")
	if err != nil {
		panic(err)
	}
	filePath := path.Join(dir, "startpage")
	f, err := os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	f.Write([]byte("foo"))
	f.Close()

	if !util.FileExists(filePath) {
		t.Errorf("Incorrect result for FileExists (false when expected true)")
	}
	if util.FileExists("/this/path/should/not/exist/on/your/computer") {
		t.Errorf("Incorrect result for FileExists (true when expected false)")
	}
}
