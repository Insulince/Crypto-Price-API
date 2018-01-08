package main

import (
	"encoding/json"
	"io/ioutil"
	"crypto-price-fetcher/pkg/models"
	"crypto-price-fetcher/pkg/routes"
	"log"
	"net/http"
	"strconv"
	"crypto-price-fetcher/pkg/database"
	"crypto-price-fetcher/pkg/engine/job"
)

var config models.Config

func main() () {
	configure()

	// TODO: Some way to resume all jobs from last run. Possibly by storing them in a json file or something.

	go job.StartEngine(config)
	//go watcher.StartEngine(config)

	router := models.CreateRouter()
	router = routes.CreateRoutes(router)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), router))
}

func configure() () {
	populateConfig()
	database.InitializeDatabase(config)
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
