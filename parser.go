package main

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

// ParseHTML parse an html content recursively and return result as a struct
func ParseHTML(content string) *html.Node {
	reader := strings.NewReader(content)
	htmldoc, err := html.Parse(reader)
	if err != nil {
		os.Exit(2)
	}
	print(htmldoc.FirstChild.Data)
	var recursiveIteratorFunc func(*html.Node)
	recursiveIteratorFunc = func(thisNode *html.Node) {
		for child := thisNode.FirstChild; child != nil; child = child.NextSibling {
			// if thisNode.Type == html.ElementNode {
			// 	println(thisNode.Data)
			// }
			println(thisNode.Data)
			recursiveIteratorFunc(child)
		}
	}

	return htmldoc
}
