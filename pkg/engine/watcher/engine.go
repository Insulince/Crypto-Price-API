package watcher

import (
	"crypto-price-fetcher/pkg/models"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"strings"
	"crypto-price-fetcher/pkg/database"
)

func StartEngine(config models.Config) () {
	for {
		Learn(func() (currencies []models.Currency) {
			body := "{\"ids\": [\"raiblocks\", \"funfair\", \"homeblockcoin\", \"ripple\", \"oyster-pearl\", \"deepbrain-chain\", \"modum\", \"bitcoin\"]}"

			response, err := http.Post("http://localhost:"+strconv.Itoa(config.Port)+"/currency", "application/json", strings.NewReader(body))
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()

			responseBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}

			type Response struct {
				Currencies []models.Currency `json:"currencies"`
			}
			var responseBodyObject Response
			err = json.Unmarshal(responseBody, &responseBodyObject)
			if err != nil {
				panic(err)
			}

			return responseBodyObject.Currencies
		}())

		time.Sleep(time.Minute * 5)
	}
}

func Learn(currencies []models.Currency) () {
	for _, currency := range currencies {
		database.CreateCurrency(currency)
	}
}
