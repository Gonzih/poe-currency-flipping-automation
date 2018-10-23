package main

import (
	"fmt"
	"log"
)

const (
	currencyListURL        = "http://currency.poe.trade/"
	currencySearchTemplate = "http://currency.poe.trade/search?league=%s&online=%s&stock=&want=%d&have=%d"
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

func (p Pair) String() string {
	return fmt.Sprintf("[%.0f %s <- %.0f %s] + [%.0f %s <- %.0f %s] = %f %s",
		p.of1.SellAmount, p.of1.SellName, p.of1.BuyAmount, p.of1.BuyName,
		p.of2.SellAmount, p.of2.SellName, p.of2.BuyAmount, p.of2.BuyName,
		p.of1.Profit(p.of2), p.of1.BuyName,
	)
}

func main() {
	// currenciesToScan := []string{"alteration", "fusing", "jeweller's", "chrome", "gcp"}
	online := true
	highestProfit := 0.0
	highestCurrency := ""

	for name := range currencyNames {
		offers1 := searchFor(online, "Delve", name, "chaos")
		offers2 := searchFor(online, "Delve", "chaos", name)

		for _, offer1 := range offers1 {
			for _, offer2 := range offers2 {
				if offer1.IsProfitable(offer2) {
					p := Pair{of1: offer1, of2: offer2}
					prof := p.of1.Profit(p.of2)
					if prof > highestProfit {
						highestProfit = prof
						highestCurrency = name
					}
					log.Println(p.String())
				}
			}
		}
	}

	log.Printf("Highest profit was with \"%s\" of %f chaos", highestCurrency, highestProfit)
}
