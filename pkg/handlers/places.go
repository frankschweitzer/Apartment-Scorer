package handlers

import (
	"fmt"
	"net/http"

	"github.com/frankschweitzer/Apartment-Scorer/pkg/services"
)

func ScoreHandler(w http.ResponseWriter, r *http.Request) {
	// Get address from query params or request body
	address := r.URL.Query().Get("address")
	score := services.CalculateApartmentScore(address)
	w.Write([]byte(fmt.Sprintf("Score for %s: %d", address, score)))
}
