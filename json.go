package scraper

import (
	"encoding/json"
)

// TagWithContent a tag and its value content
type TagWithContent struct {
	tag     string
	content string
}

// TagWithContentJSON tag and content as json object
type TagWithContentJSON struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

func toJSONfrom(lst []TagWithContent) ([]byte, error) {
	return json.Marshal(lst)
}
