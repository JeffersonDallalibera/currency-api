package service

import (
	"github.com/JeffersonDallalibera/currency-api/client"
)

func Convert(from string, to string, amount float64) (float64, float64, error) {
	rate, err := client.GetExchangeRate(from, to, amount)
	if err != nil {
		return 0, 0, err
	}

	converted := amount * rate
	return converted, rate, nil
}
