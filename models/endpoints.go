package models

type Endpoint struct {
	Convert    string
	Historical string
	Symbols    string
}

func NewEndpoints() *Endpoint {
	return &Endpoint{
		Convert:    "https://currency-converter5.p.rapidapi.com/currency/convert?format=json",
		Historical: "https://currency-converter5.p.rapidapi.com/currency/historical",
		Symbols:    "https://currency-converter5.p.rapidapi.com/currency/list",
	}
}
