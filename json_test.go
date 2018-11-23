package scraper

import (
	"testing"
)

func TestToJSONFrom(t *testing.T) {
	mockData := []TagWithContent{
		TagWithContent{Content: "a", Tag: "tag1"},
		TagWithContent{Content: "b", Tag: "tag2"},
	}

	result := toJSONfrom(mockData)
	if result == "" {
		t.Errorf("should work")
	}

	if len(result) == 0 {
		t.Errorf("invalid parsing of TagContent array")
	}
}
