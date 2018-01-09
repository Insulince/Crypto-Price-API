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
