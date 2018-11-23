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

// https://stackoverflow.com/questions/43240970/how-to-mock-http-client-do-method

func (w webURL) get() (*http.Response, error) {
	res, err := http.Get(w.address)
	print("err", err.Error())
	if err != nil {
		println("url err", err.Error())
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
