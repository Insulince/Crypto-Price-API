package models

type Currency struct {
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	CurrentPriceInUSD string `json:"current-price-in-usd"`
	CurrentPriceInBTC string `json:"current-price-in-btc"`
}
