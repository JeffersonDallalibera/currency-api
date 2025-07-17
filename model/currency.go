package model

type ConvertResponse struct {
	From      string  `json:"de"`
	To        string  `json:"para"`
	Amount    float64 `json:"valor"`
	Converted float64 `json:"convertido"`
	Rate      float64 `json:"taxa"`
}

type Currency struct {
	CurrencyCode    string `json:"code"`
	CurrencyName    string `json:"name"`
	CurrencyCountry string `json:"country"`
}
