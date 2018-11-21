package scraper

import (
	"io/ioutil"
	"net/http"
	"os"
)

func fetch(url string) string {
	res, err := http.Get(url)
	if err != nil {
		println("url err", err.Error())
		os.Exit(2)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	parsedResponse := string(body)
	return parsedResponse
}
