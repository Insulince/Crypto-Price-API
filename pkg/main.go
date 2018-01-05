package main

import (
	"encoding/json"
	"io/ioutil"
	"crypto-price-fetcher/pkg/models"
	"crypto-price-fetcher/pkg/coin-market-cap"
	"crypto-price-fetcher/pkg/twilio"
	"fmt"
)

var config models.Config

func main() () {
	configure()

	currencies := coin_market_cap.GetCurrencies([]string{"raiblocks", "bitcoin", "iota", "siacoin", "dash"})

	message := ""

	for _, currency := range currencies {
		message += currency.Name + " (" + currency.CMCRank + ") - $" + currency.CurrentPriceInUSD + "/Ƀ" + currency.CurrentPriceInBTC + " | 1h: " + currency.PercentChangeInOneHour + ", 1d: " + currency.PercentChangeInOneDay + ", 1w: " + currency.PercentChangeInOneWeek + "\n"
	}

	fmt.Printf(message)

	twilio.SendMessage(config, message)
}

func configure() () {
	populateConfig()
}

func populateConfig() () {
	jsonFile, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(jsonFile, &config)
}
