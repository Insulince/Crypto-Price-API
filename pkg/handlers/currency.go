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
		Respond(w, responses.Error{Message: "Could not process request."}, http.StatusInternalServerError)
		return
	}

	id := routeVariables["id"]

	currencyFromCMC, err := services.GetCurrency(id)
	if err != nil {
		if err.Error() != "id not found" {
			fmt.Fprintln(os.Stderr, err)
			Respond(w, responses.Error{Message: "Unexpected error encountered."}, http.StatusInternalServerError)
			return
		}
		Respond(w, responses.Error{Message: "No crypto-currency with id \"" + id + "\" was found."}, http.StatusBadRequest)
		return
	}

	type Response models.Currency
	Respond(w, Response{Name: currencyFromCMC.Name, Symbol: currencyFromCMC.Symbol, ID: currencyFromCMC.Id, CurrentPriceInUSD: currencyFromCMC.CurrentPriceInUSD, CurrentPriceInBTC: currencyFromCMC.CurrentPriceInBTC}, http.StatusOK)
}

func MultipleCurrencyFetch(w http.ResponseWriter, r *http.Request) () {
	_, _, rawPostBody, err := CallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not process request."}, http.StatusInternalServerError)
		return
	}

	type PostBody struct {
		Ids []string `json:"ids"`
	}
	var postBody PostBody
	err = json.Unmarshal(rawPostBody, &postBody)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: "Could not read request body."}, http.StatusBadRequest)
		return
	}

	currenciesFromCMC, err := services.GetCurrencies(postBody.Ids)
	if err != nil {
		if err.Error() != "id not found" {
			fmt.Fprintln(os.Stderr, err)
			Respond(w, responses.Error{Message: "Unexpected error encountered."}, http.StatusInternalServerError)
			return
		}
		Respond(w, responses.Error{Message: "One of the provided crypto-currency ids was not recognized. Provided ids were \"" + strings.Join(postBody.Ids, "\", \"") + "\"."}, http.StatusBadRequest)
		return
	}
	currencies := make([]models.Currency, len(currenciesFromCMC))
	for i, currencyFromCMC := range currenciesFromCMC {
		currencies[i] = models.Currency{Name: currencyFromCMC.Name, Symbol: currencyFromCMC.Symbol, ID: currencyFromCMC.Id, CurrentPriceInUSD: currencyFromCMC.CurrentPriceInUSD, CurrentPriceInBTC: currencyFromCMC.CurrentPriceInBTC}
	}

	type Response struct {
		Currencies []models.Currency `json:"currencies"`
	}
	Respond(w, Response{Currencies: currencies}, http.StatusOK)
}

func GlobalDataFetch(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := CallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not process request."}, http.StatusInternalServerError)
		return
	}

	globalDataFromCMC := services.GetGlobalData()

	type Response models.GlobalData
	Respond(w, Response{TotalMarketCapInUSD: globalDataFromCMC.TotalMarketCapInUSD, TotalVolumeInTwentyFourHoursInUSD: globalDataFromCMC.TotalVolumeInTwentyFourHoursInUSD}, http.StatusOK)
}
