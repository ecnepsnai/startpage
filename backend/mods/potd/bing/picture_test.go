package bing_test

import (
	"testing"

	"github.com/ecnepsnai/startpage/mods/potd/bing"
)

func TestGetPicture(t *testing.T) {
	t.Parallel()

	picture, err := bing.GetPicture()
	if err != nil {
		t.Fatalf("Error getting picture from bing: %s", err.Error())
	}
	if picture == nil {
		t.Fatalf("No picture returned when one expected")
	}
	if picture.URL() == "" {
		t.Fatalf("No picture URL returned when one expected")
	}
}
