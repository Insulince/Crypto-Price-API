package handlers

import (
	"net/http"
	"fmt"
	"os"
	"crypto-price-fetcher/pkg/models/responses"
)

func NotFound(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := CallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not process request."}, http.StatusInternalServerError)
		return
	}

	Respond(w, responses.Error{Message: "Unrecognized call."}, http.StatusNotFound)
}
