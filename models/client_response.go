package models

type CirculationResponse struct {
	BaseCurrencyCode string        `json:"base_currency_code"`
	Rates            CurrencyRates `json:"rates"`
}
