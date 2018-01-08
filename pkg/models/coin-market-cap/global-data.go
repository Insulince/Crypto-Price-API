package coin_market_cap

type GlobalData struct {
	TotalMarketCapInUSD               float64 `json:"total_market_cap_usd"`
	TotalVolumeInTwentyFourHoursInUSD float64 `json:"total_24h_volume_usd"`
	BitcoinsPercentageOfMarketCap     float64 `json:"bitcoin_percentage_of_market_cap"`
	ActiveCurrencies                  int     `json:"active_currencies"`
	ActiveAssets                      int     `json:"active_assets"`
	ActiveMarkets                     int     `json:"active_markets"`
	LastUpdatedTime                   int     `json:"last_updated"`
}
