package scraper

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

func (t TagWithContent) isValidTag(specificTag string) bool {
	isValid := t.Content != "" && t.Tag != ""

	if specificTag != "" {
		return isValid && t.Tag == specificTag
	}

	return isValid
}

func parseHTMLPage(content string, specificTag string) []TagWithContent {
	var tagsWithContent []TagWithContent
	htmldoc := getHTMLDocument(content)

	var recursiveIteratorFunc func(*html.Node)
	recursiveIteratorFunc = func(thisNode *html.Node) {
		var element TagWithContent
		if thisNode.Type == html.ElementNode {
			element.Tag = thisNode.Data
			contentChild := thisNode.FirstChild
			if contentChild != nil {
				element.Content = contentChild.Data
			}
		}
		for child := thisNode.FirstChild; child != nil; child = child.NextSibling {
			recursiveIteratorFunc(child)
		}
		if element.isValidTag(specificTag) {
			tagsWithContent = append(tagsWithContent, element)
		}
	}
	recursiveIteratorFunc(htmldoc)
	return tagsWithContent
}

func getHTMLDocument(content string) *html.Node {
	reader := strings.NewReader(content)
	htmldoc, err := html.Parse(reader)
	if err != nil {
		println("parse err", err.Error())
		os.Exit(3)
	}
	return htmldoc
}
