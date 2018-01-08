package handlers

import (
	"net/http"
	"crypto-price-fetcher/pkg/models"
)

func NotFound(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response models.Error
	Respond(w, Response{Message: "Unrecognized call."})
}
