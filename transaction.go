package main

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	Id            string    `json:"id"`
	Type          string    `json:"type"`
	Date          time.Time `json:"date"`
	Amount        float64   `json:"amount"`
	Fee           float64   `json:"fee"`
	FixedCurrency string    `json:"fixedCurrency"`
	TaxCurrency   string    `json:"taxCurrency"`
	External      string    `json:"external"`
	ExternalKey   string    `json:"externalKey"`
	Login         string    `json:"login"`
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	var err error
	type Alias Transaction
	aux := &struct {
		Date int64 `json:"date"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Date = time.Unix(aux.Date / 1000, 0)
	return nil
}
