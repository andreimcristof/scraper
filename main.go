package scraper

import (
	"flag"
	"log"
)

// ParseWithArgs parse from arguments url and tag
func ParseWithArgs(url string, tag string) string {
	return parse(url, tag)
}

// ParseWithFlags parse from flags --url and --tag
func ParseWithFlags() string {
	url, tag := getFromFlags()
	return parse(url, tag)
}

func parse(url string, tag string) string {
	response := fetch(url)
	parsed := parseHTMLPage(response, tag)
	strJSON := toJSONfrom(parsed)
	return strJSON
}

func getFromFlags() (string, string) {
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
