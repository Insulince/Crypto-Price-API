package models

type CMCCurrency struct {
	Id                           string `json:"id"`
	Name                         string `json:"name"`
	Symbol                       string `json:"symbol"`
	CMCRank                      string `json:"rank"`
	CurrentPriceInUSD            string `json:"price_usd"`
	CurrentPriceInBTC            string `json:"price_btc"`
	VolumeInTwentyFourHoursInUSD string `json:"24h_volume_usd"`
	MarketCapInUSD               string `json:"market_cap_usd"`
	AvailableSupply              string `json:"available_supply"`
	TotalSupply                  string `json:"total_supply"`
	MaximumSupply                string `json:"max_supply"`
	PercentChangeInOneHour       string `json:"percent_change_1h"`
	PercentChangeInOneDay        string `json:"percent_change_24h"`
	PercentChangeInOneWeek       string `json:"percent_change_7d"`
	LastUpdatedTime              string `json:"last_updated"`
}
