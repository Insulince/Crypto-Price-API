package coin_market_cap

import (
	"crypto-price-fetcher/pkg/models"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func GetCurrencies(currencyID []string) (currencies []models.CMCCurrency) {
	totalRequests := len(currencyID)
	requestsCompleted := 0

	requestedCurrencies := make([]models.CMCCurrency, totalRequests)

	start := time.Now()
	for _, currencyID := range currencyID {
		go func(id string) () {
			response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/" + id)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()

			responseBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}

			var cmcCurrenciesInThisResponse []models.CMCCurrency
			json.Unmarshal(responseBody, &cmcCurrenciesInThisResponse)

			requestedCurrencies[requestsCompleted] = cmcCurrenciesInThisResponse[0]

			requestsCompleted++
		}(currencyID)
	}
	for requestsCompleted < totalRequests {
		time.Sleep(10 * time.Millisecond)
	}
	elapsed := time.Since(start)

	for _, currency := range requestedCurrencies {
		fmt.Printf("%v\n", currency)
	}
	fmt.Printf("It took \"%s\" seconds to fetch all \"%v\" currencies.\n", elapsed, totalRequests)

	return requestedCurrencies
}
