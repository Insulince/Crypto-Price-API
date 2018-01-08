package handlers

import (
	"net/http"
	"crypto-price-fetcher/pkg/services"
	"encoding/json"
	"crypto-price-fetcher/pkg/models"
)

func SpecificCurrencyFetch(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, _ := CallReceived(r)

	currencyFromCMC := services.GetCurrency(routeVariables["id"])
	currency := models.Currency{Name: currencyFromCMC.Name, Symbol: currencyFromCMC.Symbol, CurrentPriceInUSD: currencyFromCMC.CurrentPriceInUSD, CurrentPriceInBTC: currencyFromCMC.CurrentPriceInBTC}

	type Response models.Currency
	Respond(w, Response{Name: currency.Name, Symbol: currency.Symbol, CurrentPriceInUSD: currency.CurrentPriceInUSD, CurrentPriceInBTC: currency.CurrentPriceInBTC})
}

func MultipleCurrencyFetch(w http.ResponseWriter, r *http.Request) () {
	_, _, postBody := CallReceived(r)
	type Body struct {
		Ids []string `json:"ids"`
	}
	var body Body
	json.Unmarshal(postBody, &body)

	currenciesFromCMC := services.GetCurrencies(body.Ids)
	currencies := make([]models.Currency, len(currenciesFromCMC))
	for i, currencyFromCMC := range currenciesFromCMC {
		currencies[i] = models.Currency{Name: currencyFromCMC.Name, Symbol: currencyFromCMC.Symbol, CurrentPriceInUSD: currencyFromCMC.CurrentPriceInUSD, CurrentPriceInBTC: currencyFromCMC.CurrentPriceInBTC}
	}

	type Response struct {
		Currencies []models.Currency `json:"currencies"`
	}
	Respond(w, Response{Currencies: currencies})
}

func GlobalDataFetch(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	globalDataFromCMC := services.GetGlobalData()

	type Response models.GlobalData
	Respond(w, Response{TotalMarketCapInUSD: globalDataFromCMC.TotalMarketCapInUSD, TotalVolumeInTwentyFourHoursInUSD: globalDataFromCMC.TotalVolumeInTwentyFourHoursInUSD})
}
