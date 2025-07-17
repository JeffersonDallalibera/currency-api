package model

type ExchangeRateResponse struct {
	Result       string  `json:"result"`        // "success"
	Documentation string `json:"documentation"` // url docs
	TermsOfUse   string  `json:"terms_of_use"`
	TimeLastUpdateUnix int64 `json:"time_last_update_unix"`
	TimeNextUpdateUnix int64 `json:"time_next_update_unix"`
	BaseCode     string  `json:"base_code"`
	TargetCode   string  `json:"target_code"`
	ConversionRate float64 `json:"conversion_rate"`
}