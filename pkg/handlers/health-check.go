package handlers

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
		Message string `json:"message"`
	}
	Respond(w, Response{Message: "OK"})
}
