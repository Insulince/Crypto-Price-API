package main

import (
	"encoding/json"
	"io/ioutil"
	"crypto-price-fetcher/pkg/models"
)

var config models.Config

func main() () {
	configure()
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
