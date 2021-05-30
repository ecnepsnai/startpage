package memoryexpress_test

import (
	"net/http"
	"testing"

	"github.com/ecnepsnai/startpage/mods/medailydeal/memoryexpress"
)

func TestGetDailyDeal(t *testing.T) {
	t.Parallel()

	resp, err := http.Get("https://www.memoryexpress.com/")
	if err != nil {
		t.Fatalf("Error getting memory express index: %s", err.Error())
	}
	defer resp.Body.Close()
	deal, err := memoryexpress.GetDailyDeal(resp.Body)
	if err != nil {
		t.Fatalf("Error getting daily deal: %s", err.Error())
	}
	if deal == nil {
		t.Fatalf("No deal returned")
	}
	t.Logf("Deal: %#v\n", *deal)
}
