package main

import (
	"github.com/shopspring/decimal"
)

type Orderbook struct {
	Asks [][]decimal.Decimal `json:"asks"`
	Bids [][]decimal.Decimal `json:"bids"`
}
