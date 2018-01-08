package main

import (
	"encoding/json"
	"io/ioutil"
	"crypto-price-fetcher/pkg/models"
	"crypto-price-fetcher/pkg/routes"
	"log"
	"net/http"
	"strconv"
)

var config models.Config

func main() () {
	configure()

	//currencies := services.GetCurrencies([]string{"raiblocks", "bitcoin", "iota", "siacoin", "dash"})
	//message := ""
	//for _, currency := range currencies {
	//	message += currency.Name + " (" + currency.CoinMarketCapRank + ") - $" + currency.CurrentPriceInUSD + "/Ƀ" + currency.CurrentPriceInBTC + " | 1h: " + currency.PercentChangeInOneHour + ", 1d: " + currency.PercentChangeInOneDay + ", 1w: " + currency.PercentChangeInOneWeek + "\n"
	//}
	//fmt.Printf(message)

	// TODO: Some way to resume all jobs from last run. Possibly by storing them in a json file or something.
	router := models.CreateRouter()
	router = routes.CreateRoutes(router)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), router))
}

func configure() () {
	populateConfig()
}

func populateConfig() () {
	jsonFile, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		panic(err)
	}
}
