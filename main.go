package main

import (
	"fmt"
	"log"
	"strings"
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
	of  SearchOffer
	ofs []SearchOffer
}

func (p Pair) String() string {
	buf := strings.Builder{}
	buf.WriteString(p.of.String())
	buf.WriteString("\n")

	for _, offer := range p.ofs {
		buf.WriteString("\t -> ")
		buf.WriteString(offer.String())
		buf.WriteString(fmt.Sprintf(" %f", p.Profit(offer)))
		buf.WriteString("\n")
	}

	return buf.String()
}

func (p Pair) Profit(of2 SearchOffer) float64 {
	return p.of.Profit(of2)
}

func (p Pair) MaxProfit() float64 {
	profit := 0.0

	for _, offer := range p.ofs {
		np := p.Profit(offer)
		if np > profit {
			profit = np
		}
	}

	return profit
}

func main() {
	Execute()
}
