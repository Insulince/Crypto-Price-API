package routes

import (
	"crypto-price-fetcher/pkg/handlers"
	"crypto-price-fetcher/pkg/models"
	"net/http"
)

func CreateRoutes(router *models.Router) (*models.Router) {
	router.HandleFunc("/", handlers.Home).Methods("GET")

	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	router.HandleFunc("/currency/global", handlers.GlobalDataFetch).Methods("GET")
	router.HandleFunc("/currency/{id}", handlers.SpecificCurrencyFetch).Methods("GET")
	router.HandleFunc("/currency", handlers.MultipleCurrencyFetch).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	return router
}
