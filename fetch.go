package scraper

import (
	"io/ioutil"
	"net/http"
	"os"
)

type httpClient interface {
	get() (*http.Response, error)
}

type webURL struct {
	address string
}

func (w webURL) get() (*http.Response, error) {
	res, err := http.Get(w.address)
	if err != nil {
		println("httpGet err", err.Error())
		os.Exit(2)
	}
	defer res.Body.Close()
	return res, err
}

func fetch(address string) string {
	wu := webURL{address: address}
	res, _ := wu.get()
	body, _ := ioutil.ReadAll(res.Body)
	parsedResponse := string(body)
	return parsedResponse
}
