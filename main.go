package main

import (
	"log"
)

const (
	currencyListURL        = "http://currency.poe.trade/"
	currencySearchTemplate = "http://currency.poe.trade/search?league=Delve&online=x&stock=&want=%d&have=%d"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Pair struct {
	of1 SearchOffer
	of2 SearchOffer
}

func main() {
	log.Println(currencyNames)

	offers1 := searchFor("fusing", "chaos")
	offers2 := searchFor("chaos", "fusing")

	for _, offer1 := range offers1 {
		for _, offer2 := range offers2 {
			if offer1.IsProfitable(offer2) {
				p := Pair{of1: offer1, of2: offer2}
				log.Println(p)
				log.Println(p.of1.Profit(p.of2))
			}
		}
	}
}
