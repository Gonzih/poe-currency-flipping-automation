package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

var currencyNames map[string]int64

func lookupCurrencyByName(name string) int64 {
	id, ok := currencyNames[name]
	if !ok {
		log.Fatalf("Unknown currency %s", name)
	}

	return id
}

func lookupCurrencyByID(searchID int64) string {
	for name, id := range currencyNames {
		if id == searchID {
			return name
		}
	}
	return ""
}

func initCurrencyNames() {
	log.Println("Initializing currency table")
	currencyNames = make(map[string]int64, 0)
	reader, err := GET(currencyListURL)
	must(err)
	defer reader.Close()
	document, err := goquery.NewDocumentFromReader(reader)
	must(err)
	document.Find("div.currency-selectable.currency-square ").Each(func(index int, element *goquery.Selection) {
		title, _ := element.Attr("data-title")
		id, _ := element.Attr("data-id")
		must(err)
		currencyNames[title] = s2i(id)
	})
}

func init() {
	initCurrencyNames()
}
