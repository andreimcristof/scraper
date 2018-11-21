package scraper

import (
	"flag"
	"log"
)

// ParseWithArgs parse from arguments url and tag
func ParseWithArgs(url string, tag string) ([]byte, error) {
	tagsWithContent := parse(url, tag)
	return toJSONfrom(tagsWithContent)
}

// ParseWithFlags parse from flags --url and --tag
func ParseWithFlags() ([]byte, error) {
	url, tag := setupFlags()
	tagsWithContent := parse(url, tag)
	return toJSONfrom(tagsWithContent)
}

func parse(url string, tag string) []TagWithContent {
	response := fetch(url)
	parsed := parseHTMLPage(response, tag)
	return parsed
}

func setupFlags() (string, string) {
	urlArgHint := "enter url to parse"
	tagArgHint := "(optional) enter a specific tag to retrieve only"
	urlFlagPtr := flag.String("url", "", urlArgHint)
	tagFlagPtr := flag.String("tag", "", tagArgHint)
	flag.Parse()
	if *urlFlagPtr == "" {
		log.Fatal("must specify --url flag")
	}
	println("parsing url:", *urlFlagPtr, "| tag:", *tagFlagPtr)
	return *urlFlagPtr, *tagFlagPtr
}
