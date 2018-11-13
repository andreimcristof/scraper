package htmlparser

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

// ParseHTML parse an html content recursively and return result as a struct
func ParseHTML(content string) *html.Node {
	reader := strings.NewReader(content)
	parser, err := html.Parse(reader)
	if err != nil {
		os.Exit(2)
	}

	return parser
}
