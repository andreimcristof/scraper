package scraper

import (
	"flag"
	"log"
)

// ScrapeWithArgs parse from arguments url and tag
func ScrapeWithArgs(url string, tag string) string {
	return scrape(url, tag)
}

// ScrapeWithFlags parse from flags --url and --tag
func ScrapeWithFlags() string {
	url, tag := getFromFlags()
	return scrape(url, tag)
}

func scrape(url string, tag string) string {
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
