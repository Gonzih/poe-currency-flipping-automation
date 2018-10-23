package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func searchFor(online bool, league, want, have string) []SearchOffer {
	wantID, ok := currencyNames[want]
	if !ok {
		log.Fatalf("Unknown currency: %s", want)
	}

	haveID, ok := currencyNames[have]
	if !ok {
		log.Fatalf("Unknown currency: %s", have)
	}

	onlineFilter := ""
	if online {
		onlineFilter = "x"
	}

	url := fmt.Sprintf(currencySearchTemplate, league, onlineFilter, wantID, haveID)
	reader, err := GET(url)
	must(err)
	defer reader.Close()
	document, err := goquery.NewDocumentFromReader(reader)
	must(err)

	offers := make([]SearchOffer, 0)

	document.Find("div.displayoffer").Each(func(index int, element *goquery.Selection) {
		sellAmount, _ := element.Attr("data-sellvalue")
		sellCurrency, _ := element.Attr("data-sellcurrency")
		sellName := lookupCurrencyByID(s2i(sellCurrency))

		buyAmount, _ := element.Attr("data-buyvalue")
		buyCurrency, _ := element.Attr("data-buycurrency")
		buyName := lookupCurrencyByID(s2i(buyCurrency))

		ign, _ := element.Attr("data-ign")
		stock, _ := element.Attr("data-stock")

		offer := NewSearchOffer(
			sellAmount, sellName,
			buyAmount, buyName,
			ign, stock)
		offers = append(offers, offer)
	})

	return offers
}
