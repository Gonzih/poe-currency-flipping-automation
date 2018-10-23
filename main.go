package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const (
	currencyListURL        = "http://currency.poe.trade/"
	currencySearchTemplate = "http://currency.poe.trade/search?league=Delve&online=x&stock=&want=%d&have=%d"
)

func GET(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	return response.Body, err
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func s2i(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Error parsing %s to int", s)
	}

	return i
}

func s2f(s string) float64 {
	i, err := strconv.ParseFloat(s, 10)
	if err != nil {
		log.Fatalf("Error parsing %s to float", s)
	}

	return i
}

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

func main() {
	log.Println(currencyNames)

	log.Println("=================")
	offers := searchFor("fusing", "chaos")
	for i, offer := range offers {
		log.Println(offer.String())
		if i == 5 {
			break
		}
	}

	log.Println("=================")
	offers = searchFor("chaos", "fusing")
	for i, offer := range offers {
		log.Println(offer.String())
		if i == 5 {
			break
		}
	}
	log.Println("=================")
}

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

func searchFor(want, have string) []SearchOffer {
	wantID, ok := currencyNames[want]
	if !ok {
		log.Fatalf("Unknown currency: %s", want)
	}

	haveID, ok := currencyNames[have]
	if !ok {
		log.Fatalf("Unknown currency: %s", have)
	}

	url := fmt.Sprintf(currencySearchTemplate, wantID, haveID)
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
