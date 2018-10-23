package main

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

const cacheFilePath = "/tmp/poe-currencies.bin"

var currencyNames map[string]int64

func saveCurrencyTable() error {
	file, err := os.Create(cacheFilePath)
	defer file.Close()
	if err != nil {
		return err
	}

	enc := gob.NewEncoder(file)
	return enc.Encode(currencyNames)
}

func loadCurrencyTable() error {
	file, err := os.Open(cacheFilePath)
	defer file.Close()
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(file)
	return dec.Decode(&currencyNames)
}

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
	err := loadCurrencyTable()
	if err != nil {
		initCurrencyNames()
		err = saveCurrencyTable()
		if err != nil {
			log.Fatalf("Error saving currency table: %s", err)
		}
	}
}
