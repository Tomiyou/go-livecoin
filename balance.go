package main

type Balance struct {
	Type     string  `json:"type"`
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}
