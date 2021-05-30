package potd_test

import (
	"os"
	"testing"

	"github.com/ecnepsnai/logtic"
	"github.com/ecnepsnai/startpage/mods/potd"
)

func TestMain(m *testing.M) {
	logtic.Log.Level = logtic.LevelDebug
	logtic.Open()
	os.Exit(m.Run())
}

func TestPotd(t *testing.T) {
	t.Parallel()

	tmpDir, err := os.MkdirTemp("", "potd")
	if err != nil {
		panic(err)
	}
	i, err := potd.Setup(tmpDir, nil)
	if err != nil {
		panic(err)
	}

	if err := i.Refresh(); err != nil {
		t.Fatalf("Error refreshing: %s", err.Error())
	}
	if err := i.Refresh(); err != nil {
		t.Fatalf("Error refreshing: %s", err.Error())
	}
}
