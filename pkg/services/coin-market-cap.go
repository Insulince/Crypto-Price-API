package services

import (
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"crypto-price-fetcher/pkg/models/coin-market-cap"
	"errors"
)

func GetCurrency(currencyID string) (currency coin_market_cap.Currency, err error) {
	response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/" + currencyID)
	if err != nil {
		return coin_market_cap.Currency{}, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return coin_market_cap.Currency{}, err
	}

	var currenciesInThisResponse []coin_market_cap.Currency
	err = json.Unmarshal(responseBody, &currenciesInThisResponse)
	if err != nil {
		var CMCError coin_market_cap.Error
		err = json.Unmarshal(responseBody, &CMCError)
		if err != nil {
			panic(err)
		}

		if CMCError.Message != "id not found" {
			panic(errors.New("Unknown error: " + CMCError.Message))
		}
		return coin_market_cap.Currency{}, errors.New("id not found")
	}

	return currenciesInThisResponse[0], nil
}

func GetCurrencies(currencyIDs []string) (currencies []coin_market_cap.Currency, err error) {
	totalRequests := len(currencyIDs)
	requestsCompleted := 0

	requestedCurrencies := make([]coin_market_cap.Currency, totalRequests)
	currencyChannel := make(chan coin_market_cap.Currency)
	errChannel := make(chan error)

	for _, currencyID := range currencyIDs {
		go func(currencyID string, currencyChannel chan coin_market_cap.Currency) () {
			currency, err := GetCurrency(currencyID)
			if err != nil {
				errChannel <- err
			}
			currencyChannel <- currency
		}(currencyID, currencyChannel)
	}
	for requestsCompleted < totalRequests {
		select {
		case err := <-errChannel:
			return nil, err
		case currency := <-currencyChannel:
			requestedCurrencies[requestsCompleted] = currency
			requestsCompleted++
		}
		time.Sleep(10 * time.Millisecond)
	}

	return requestedCurrencies, nil
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

	err = json.Unmarshal(responseBody, &globalData)
	if err != nil {
		panic(err)
	}

	return globalData
}
