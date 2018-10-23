package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOfferBasic(t *testing.T) {
	offer := NewSearchOffer("100.0", "fusing", "33.0", "chaos", "potato", "")
	assert.Equal(t, offer.SellAmount, 100.0)
	assert.Equal(t, offer.BuyAmount, 33.0)
}

func TestOfferRate(t *testing.T) {
	offer := NewSearchOffer("100.0", "fusing", "33.0", "chaos", "potato", "")
	assert.Equal(t, 0.33, offer.RateOf("chaos"))
	assert.Equal(t, 3.0303030303030303, offer.RateOf("fusing"))
}

func TestOfferCompatible(t *testing.T) {
	offer1 := NewSearchOffer("100.0", "chaos", "20.0", "fusing", "potato", "")
	offer2 := NewSearchOffer("100.0", "fusing", "20.0", "chaos", "potato", "")
	assert.True(t, offer1.IsCompatible(offer2))
}

func TestOfferProfitableFalse(t *testing.T) {
	offer1 := NewSearchOffer("100.0", "fusing", "33.0", "chaos", "potato", "")
	offer2 := NewSearchOffer("20.0", "chaos", "61.0", "fusing", "potato", "")
	assert.False(t, offer1.IsProfitable(offer2))
}

func TestOfferProfitableTrue(t *testing.T) {
	offer1 := NewSearchOffer("100.0", "fusing", "33.0", "chaos", "potato", "")
	offer2 := NewSearchOffer("30.0", "chaos", "90.0", "fusing", "potato", "")
	assert.True(t, offer1.IsProfitable(offer2))
}

func TestOfferProfit(t *testing.T) {
	offer1 := NewSearchOffer("100.0", "fusing", "33.0", "chaos", "potato", "")
	offer2 := NewSearchOffer("30.0", "chaos", "40.0", "fusing", "potato", "")
	assert.Equal(t, 0.42, offer1.Profit(offer2))
}
