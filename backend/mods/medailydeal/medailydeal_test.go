package medailydeal_test

import (
	"os"
	"testing"

	"github.com/ecnepsnai/logtic"
	"github.com/ecnepsnai/startpage/mods/medailydeal"
)

func TestMain(m *testing.M) {
	logtic.Log.Level = logtic.LevelDebug
	logtic.Open()
	os.Exit(m.Run())
}

func TestMedailydeal(t *testing.T) {
	t.Parallel()

	tmpDir, err := os.MkdirTemp("", "medailydeal")
	if err != nil {
		panic(err)
	}
	i, err := medailydeal.Setup(tmpDir, nil)
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
