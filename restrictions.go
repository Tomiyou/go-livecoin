package main

import (
	"github.com/shopspring/decimal"
)

type Restrictions struct {
	Success      bool            `json:"success"`
	MinBtcVolume decimal.Decimal `json:"minBtcVolume"`
	Restrictions []Restriction   `json:"restrictions"`
}

type Restriction struct {
	CurrencyPair string `json:"currencyPair"`
	PriceScale   int    `json:"priceScale"`
}
