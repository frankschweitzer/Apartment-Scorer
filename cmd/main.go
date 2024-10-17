package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/frankschweitzer/Apartment-Scorer/pkg/services"
)

// Handler to fetch nearby places
func fetchPlacesHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query params
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		return
	}

	category := r.URL.Query().Get("category")
	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	// Convert address to coordinates
	lat, lon, err := services.FetchLatLonFromAddress(address)
	if err != nil {
		http.Error(w, "Error locating coordinates", http.StatusInternalServerError)
		return
	}

	// Use the address to convert it to lat and lon to get a circular filter
	radiusMeters := 500
	filter := fmt.Sprintf("circle:%f,%f,%d", lon, lat, radiusMeters)

	// Call the FetchNearbyPlaces service
	places, err := services.FetchNearbyRestaurants(category, filter)
	if err != nil {
		http.Error(w, "Error fetching places", http.StatusInternalServerError)
		return
	}

	// Set response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Convert the places to JSON and write to the response
	json.NewEncoder(w).Encode(places)
}

func main() {
	// Define routes
	http.HandleFunc("/nearby-places", fetchPlacesHandler)

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
