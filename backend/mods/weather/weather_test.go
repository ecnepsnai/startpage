package weather_test

import (
	"os"
	"testing"

	"github.com/ecnepsnai/logtic"
	"github.com/ecnepsnai/startpage/mods/weather"
)

func TestMain(m *testing.M) {
	logtic.Log.Level = logtic.LevelDebug
	logtic.Open()
	os.Exit(m.Run())
}

func TestRefresh(t *testing.T) {
	t.Parallel()

	if os.Getenv("OW_API_KEY") == "" {
		t.Skip("OW_API_KEY environment variable not defined, skipping test")
	}

	tmpDir, err := os.MkdirTemp("", "potd")
	if err != nil {
		panic(err)
	}

	i, err := weather.Setup(tmpDir, &weather.Options{
		Latitude:  49.2829766,
		Longitude: -123.1204358,
		APIKey:    os.Getenv("OW_API_KEY"),
	})
	if err != nil {
		t.Fatalf("Error setting up instance: %s", err.Error())
	}

	if err := i.Refresh(); err != nil {
		t.Fatalf("Error refreshing data: %s", err.Error())
	}
	if err := i.Refresh(); err != nil {
		t.Fatalf("Error refreshing data: %s", err.Error())
	}
}
