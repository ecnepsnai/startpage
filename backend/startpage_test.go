package startpage

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"testing"

	"github.com/ecnepsnai/logtic"
)

var tmpDir string

func testSetup() {
	t, err := os.MkdirTemp("", "startpage")
	if err != nil {
		panic(err)
	}
	tmpDir = t
	dataDirectory = tmpDir

	for _, arg := range os.Args {
		if arg == "-test.v=true" {
			logtic.Log.Level = logtic.LevelDebug
			logtic.Open()
		}
	}

	fsSetup()
	LoadOptions()
	dataStoreSetup()
}

func testTeardown() {
	dataStoreTeardown()
	os.RemoveAll(tmpDir)
}

func TestMain(m *testing.M) {
	testSetup()
	r := m.Run()
	testTeardown()
	os.Exit(r)
}

func randomString(length uint16) string {
	randB := make([]byte, length)
	if _, err := rand.Read(randB); err != nil {
		panic(err)
	}
	return hex.EncodeToString(randB)
}
