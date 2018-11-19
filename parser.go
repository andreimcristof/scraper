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

func (t TagWithContent) isValidTag() bool {
	return t.content != "" && t.tag != ""
}

// ParseHTML parse an html content recursively and return result as a struct
//TODO: export as pkg somehow
func ParseHTML(content string) []TagWithContent {
	var tagsWithContent []TagWithContent

	reader := strings.NewReader(content)
	htmldoc, err := html.Parse(reader)
	if err != nil {
		os.Exit(2)
	}

	var recursiveIteratorFunc func(*html.Node)
	recursiveIteratorFunc = func(thisNode *html.Node) {
		var element TagWithContent
		if thisNode.Type == html.ElementNode {
			element.tag = thisNode.Data
			contentChild := thisNode.FirstChild
			if contentChild != nil {
				element.content = contentChild.Data
			}
		}
		for child := thisNode.FirstChild; child != nil; child = child.NextSibling {
			recursiveIteratorFunc(child)
		}
		if element.isValidTag() {

			tagsWithContent = append(tagsWithContent, element)
		}
	}

	recursiveIteratorFunc(htmldoc)
	return tagsWithContent
}
