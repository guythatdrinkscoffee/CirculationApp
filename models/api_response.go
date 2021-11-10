package models

import (
	"encoding/json"
	"io"
)

type CurrencyRates map[string]Rate

type APIResponse struct {
	BaseCurrencyCode string        `json:"base_currency_code"`
	BaseCurrencyName string        `json:"base_currency_name"`
	Amount           string        `json:"amount"`
	UpdatedDate      string        `json:"updated_date"`
	Rates            CurrencyRates `json:"rates"`
	Status           string        `json:"status"`
}

type Rate struct {
	CurrencyName  string `json:"currency_name"`
	Rate          string `json:"rate"`
	RateForAmount string `json:"rate_for_amount"`
}

type APIErrorResponse struct {
	Status string   `json:"status"`
	Result APIError `json:"error"`
}

func (a *APIErrorResponse) Error() string {
	return a.Result.Message
}

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (a *APIResponse) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(a)
}
