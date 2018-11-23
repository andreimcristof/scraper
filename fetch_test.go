package scraper

import (
	"testing"
)

func TestFetch(t *testing.T) {
	result := fetch("https://google.se")
	if result == "" {
		t.Errorf("should work to make http request")
	}
}
