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
	return fmt.Sprintf("[%.0f %s <- %.0f %s] by %s",
		of.SellAmount, of.SellName, of.BuyAmount, of.BuyName, of.IGN)
}
func (of SearchOffer) ToMessage(league string) string {
	if league == "" {
		league = "Standard"
	}

	return fmt.Sprintf("@%s Hi, I'd like to buy your %.0f %s for my %.0f %s in %s.",
		of.IGN, of.SellAmount, of.SellName, of.BuyAmount, of.BuyName, league)
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
