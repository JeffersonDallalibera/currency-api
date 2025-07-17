package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JeffersonDallalibera/currency-api/model"

	"github.com/JeffersonDallalibera/currency-api/config"
)



func GetExchangeRate(from, to string, amount float64) (float64, error) {
	var url string

	if amount <= 0 {
		url = fmt.Sprintf("%s/%s/pair/%s/%s", config.APIBaseURL, config.APIKey, from, to)
	} else {
		url = fmt.Sprintf("%s/%s/pair/%s/%s/%f", config.APIBaseURL, config.APIKey, from, to, amount)
	}
	
	
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result model.ExchangeRateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	if result.Result != "success" {
		return 0, fmt.Errorf("Erro na API: %v", result.Result)
	}

	return result.ConversionRate, nil
}
