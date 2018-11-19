package main

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

// TagWithContent a tag and its value content
type TagWithContent struct {
	tag     string
	content string
}

// ParseHTML parse an html content recursively and return result as a struct
//TODO: export as pkg somehow
func ParseHTML(content string) []TagWithContent {
	reader := strings.NewReader(content)
	htmldoc, err := html.Parse(reader)
	if err != nil {
		os.Exit(2)
	}

	var recursiveIteratorFunc func(*html.Node)
	recursiveIteratorFunc = func(thisNode *html.Node) {
		if thisNode.Type == html.ElementNode {
			t := thisNode.FirstChild
			if t != nil {
				println(t.Data)
			}
		}
		for child := thisNode.FirstChild; child != nil; child = child.NextSibling {
			recursiveIteratorFunc(child)
		}
	}

	recursiveIteratorFunc(htmldoc)
}
