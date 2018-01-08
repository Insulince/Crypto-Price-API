package handlers

import (
	"net/http"
	"crypto-price-fetcher/pkg/services"
	"encoding/json"
	"crypto-price-fetcher/pkg/models"
	"strings"
)

func SpecificCurrencyFetch(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, _ := CallReceived(r)
	id := routeVariables["id"]

	currencyFromCMC, err := services.GetCurrency(id)
	if err != nil {
		if err.Error() != "id not found" {
			panic(err)
		}
		type Response models.Error
		Respond(w, Response{Message: "No crypto-currency with id \"" + id + "\" was found."})
		return
	}

	type Response models.Currency
	Respond(w, Response{Name: currencyFromCMC.Name, Symbol: currencyFromCMC.Symbol, ID: currencyFromCMC.Id, CurrentPriceInUSD: currencyFromCMC.CurrentPriceInUSD, CurrentPriceInBTC: currencyFromCMC.CurrentPriceInBTC})
}

func MultipleCurrencyFetch(w http.ResponseWriter, r *http.Request) () {
	_, _, postBody := CallReceived(r)
	type Body struct {
		Ids []string `json:"ids"`
	}
	var body Body
	err := json.Unmarshal(postBody, &body)
	if err != nil {
		panic(err)
	}

	currenciesFromCMC, err := services.GetCurrencies(body.Ids)
	if err != nil {
		if err.Error() != "id not found" {
			panic(err)
		}
		type Response models.Error
		Respond(w, Response{Message: "One of the provided crypto-currency ids was not recognized. Provided ids were \"" + strings.Join(body.Ids, "\", \"") + "\"."})
		return
	}
	currencies := make([]models.Currency, len(currenciesFromCMC))
	for i, currencyFromCMC := range currenciesFromCMC {
		currencies[i] = models.Currency{Name: currencyFromCMC.Name, Symbol: currencyFromCMC.Symbol, ID: currencyFromCMC.Id, CurrentPriceInUSD: currencyFromCMC.CurrentPriceInUSD, CurrentPriceInBTC: currencyFromCMC.CurrentPriceInBTC}
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
