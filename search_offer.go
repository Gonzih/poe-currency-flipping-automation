package main

import (
	"fmt"
	"log"
)

type SearchOffer struct {
	SellAmount   float64
	SellCurrency int64
	SellName     string
	BuyAmount    float64
	BuyCurrency  int64
	BuyName      string
	IGN          string
	Stock        int64
}

func (of SearchOffer) String() string {
	return fmt.Sprintf("%s %.0f <- %.0f %s (%.5f <- %.5f) (%.5f <- %.5f)", of.SellName, of.SellAmount, of.BuyAmount, of.BuyName,
		of.SellAmount/of.SellAmount, of.BuyAmount/of.SellAmount,
		of.SellAmount/of.BuyAmount, of.BuyAmount/of.BuyAmount)
}

func (of1 SearchOffer) IsCompatible(of2 SearchOffer) bool {
	return of1.SellCurrency == of2.BuyCurrency && of1.BuyCurrency == of2.SellCurrency
}

func (of1 SearchOffer) IsProfitable(of2 SearchOffer) bool {
	return of1.IsCompatible(of2) &&
		of1.RateOf(of1.BuyName) < of2.RateOf(of1.BuyName)
}

func (of1 SearchOffer) Profit(of2 SearchOffer) float64 {
	return of2.RateOf(of1.BuyName) - of1.RateOf(of1.BuyName)
}

func (of SearchOffer) RateOf(name string) float64 {
	if name == of.BuyName {
		return of.BuyAmount / of.SellAmount
	}
	if name == of.SellName {
		return of.SellAmount / of.BuyAmount
	}

	log.Printf("Offer has no currency %s", name)
	return 0.0
}

func NewSearchOffer(sAmount, sName, bAmount, bName, IGN, stock string) SearchOffer {
	if stock == "" {
		stock = "0"
	}

	return SearchOffer{
		SellAmount:   s2f(sAmount),
		SellCurrency: currencyNames[sName],
		SellName:     sName,

		BuyAmount:   s2f(bAmount),
		BuyCurrency: currencyNames[bName],
		BuyName:     bName,

		IGN:   IGN,
		Stock: s2i(stock),
	}
}
