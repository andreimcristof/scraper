package scraper

import (
	"encoding/json"
)

// TagWithContent a tag and its value content
type TagWithContent struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

func toJSONfrom(lst []TagWithContent) string {
	marshalled, _ := json.Marshal(lst)
	jsonStrResult := string(marshalled)
	return jsonStrResult
}
