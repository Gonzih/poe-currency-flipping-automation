package main

import (
	"fmt"
	"log"
	"sort"
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
		p.Profit(), p.of1.BuyName,
	)
}

func (p Pair) Profit() float64 {
	return p.of1.Profit(p.of2)
}

func main() {
	// currenciesToScan := []string{"alteration", "fusing", "jeweller's", "chrome", "gcp"}
	online := true

	pairs := make([]Pair, 0)

	for name := range currencyNames {
		// for _, name := range currenciesToScan {
		switch name {
		case "wisdom", "chaos":
			contunue
		default:
		}

		offers1 := searchFor(online, "Delve", name, "chaos")
		offers2 := searchFor(online, "Delve", "chaos", name)

		for _, offer1 := range offers1 {
			for _, offer2 := range offers2 {
				if offer1.IsProfitable(offer2) {
					p := Pair{of1: offer1, of2: offer2}
					// log.Println(p.String())
					pairs = append(pairs, p)
				}
			}
		}
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Profit() > pairs[j].Profit() })

	for i := 0; i < 15; i++ {
		if i > len(pairs) {
			break
		}
		log.Println(pairs[i].String())
	}
}
