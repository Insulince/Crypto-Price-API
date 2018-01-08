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

	router.HandleFunc("/job", handlers.GetJobs).Methods("GET")
	router.HandleFunc("/job", handlers.CreateJob).Methods("POST")
	router.HandleFunc("/job/{id}", handlers.GetJob).Methods("GET")
	router.HandleFunc("/job/{id}", handlers.UpdateJob).Methods("PUT")
	router.HandleFunc("/job/{id}", handlers.DeleteJob).Methods("DELETE")
	router.HandleFunc("/job/start", handlers.StartProvidedJobs).Methods("POST")
	router.HandleFunc("/job/stop", handlers.StopAllJobs).Methods("GET")
	router.HandleFunc("/job/start/{id}", handlers.StartJob).Methods("GET")
	router.HandleFunc("/job/stop/{id}", handlers.StopJob).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	return router
}
