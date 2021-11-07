package models

import (
	"encoding/json"
	"io"
)

type APIResponse struct {
	BaseCurrencyCode string        `json:"base_currency_code"`
	BaseCurrencyName string        `json:"base_currency_name"`
	Amount           string        `json:"amount"`
	UpdatedDate      string        `json:"updated_date"`
	Rates            CurrencyRates `json:"rates"`
	Status           string        `json:"status"`
}

type CurrencyRates map[string]Rate

type Rate struct {
	CurrencyName  string `json:"currency_name"`
	Rate          string `json:"rate"`
	RateForAmount string `json:"rate_for_amount"`
}

func (a *APIResponse) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(a)
}
