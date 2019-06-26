package main

import (
	"encoding/json"
	"time"
)

type Trade struct {
	Id            uint64    `json:"id"`
	ClientOrderId uint64    `json:"clientorderid"`
	Type          string    `json:"type"`
	Symbol        string    `json:"symbol"`
	Date          time.Time `json:"datetime"`
	Price         float64   `json:"price"`
	Quantity      float64   `json:"quantity"`
	Commission    float64   `json:"commission"`
}

func (t *Trade) UnmarshalJSON(data []byte) error {
	var err error
	type Alias Trade
	aux := &struct {
		Date int64 `json:"datetime"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Date = time.Unix(aux.Date, 0)
	return nil
}
