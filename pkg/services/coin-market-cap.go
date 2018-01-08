package services

import (
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"crypto-price-fetcher/pkg/models/coin-market-cap"
)

func GetCurrency(currencyID string) (currency coin_market_cap.Currency) {
	response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/" + currencyID)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var currenciesInThisResponse []coin_market_cap.Currency
	json.Unmarshal(responseBody, &currenciesInThisResponse)

	return currenciesInThisResponse[0]
}

func GetCurrencies(currencyIDs []string) (currencies []coin_market_cap.Currency) {
	totalRequests := len(currencyIDs)
	requestsCompleted := 0

	requestedCurrencies := make([]coin_market_cap.Currency, totalRequests)
	currencyChannel := make(chan coin_market_cap.Currency)

	for _, currencyID := range currencyIDs {
		go func(currencyID string, currencyChannel chan coin_market_cap.Currency) {
			currencyChannel <- GetCurrency(currencyID)
		}(currencyID, currencyChannel)
	}
	for requestsCompleted < totalRequests {
		select {
		case currency := <-currencyChannel:
			requestedCurrencies[requestsCompleted] = currency
			requestsCompleted++
		}
		time.Sleep(10 * time.Millisecond)
	}

	return requestedCurrencies
}

func GetGlobalData() (globalData coin_market_cap.GlobalData) {
	response, err := http.Get("https://api.coinmarketcap.com/v1/global")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(responseBody, &globalData)

	return globalData
}
