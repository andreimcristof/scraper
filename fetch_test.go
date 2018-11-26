package scraper

import (
	"testing"
)

func TestFetch(t *testing.T) {
	t.Skip()
	result := fetch("https://google.se")
	if result == "" {
		// find some meaningful test here and inject http response in client
	}
}
