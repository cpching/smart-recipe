package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	// "os"
)

func main() {
	urls := []string{
		"",
	}
	for _, url := range urls {
		fetchPrices(url)
	}
}

func fetchPrices(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("")

	// TODO:Insert or update the DB
}
