package handlers

import (
	"net/http"
	"crypto-price-fetcher/pkg/services"
	"encoding/json"
	"crypto-price-fetcher/pkg/models"
	"strings"
	"fmt"
	"os"
	"crypto-price-fetcher/pkg/models/responses"
)

func SpecificCurrencyFetch(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, _, err := CallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not process request."})
		return
	}

	id := routeVariables["id"]

	currencyFromCMC, err := services.GetCurrency(id)
	if err != nil {
		if err.Error() != "id not found" {
			fmt.Fprintln(os.Stderr, err)
			Respond(w, responses.Error{Message: "Unexpected error encountered."})
			return
		}
		Respond(w, responses.Error{Message: "No crypto-currency with id \"" + id + "\" was found."})
		return
	}

	type Response models.Currency
	Respond(w, Response{Name: currencyFromCMC.Name, Symbol: currencyFromCMC.Symbol, ID: currencyFromCMC.Id, CurrentPriceInUSD: currencyFromCMC.CurrentPriceInUSD, CurrentPriceInBTC: currencyFromCMC.CurrentPriceInBTC})
}

func MultipleCurrencyFetch(w http.ResponseWriter, r *http.Request) () {
	_, _, rawPostBody, err := CallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not process request."})
		return
	}

	type PostBody struct {
		Ids []string `json:"ids"`
	}
	var postBody PostBody
	err = json.Unmarshal(rawPostBody, &postBody)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: "Could not read request body."})
		return
	}

	currenciesFromCMC, err := services.GetCurrencies(postBody.Ids)
	if err != nil {
		if err.Error() != "id not found" {
			fmt.Fprintln(os.Stderr, err)
			Respond(w, responses.Error{Message: "Unexpected error encountered."})
			return
		}
		Respond(w, responses.Error{Message: "One of the provided crypto-currency ids was not recognized. Provided ids were \"" + strings.Join(postBody.Ids, "\", \"") + "\"."})
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
	_, _, _, err := CallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not process request."})
		return
	}

	globalDataFromCMC := services.GetGlobalData()

	type Response models.GlobalData
	Respond(w, Response{TotalMarketCapInUSD: globalDataFromCMC.TotalMarketCapInUSD, TotalVolumeInTwentyFourHoursInUSD: globalDataFromCMC.TotalVolumeInTwentyFourHoursInUSD})
}
