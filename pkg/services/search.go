package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/frankschweitzer/Apartment-Scorer/pkg/models"
	"github.com/joho/godotenv"
)

// Sample function to search for restaurants near an address
func FetchNearbyRestaurants(category, filter string) ([]models.Place, error) {
	// Load the API key from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	apiKey := os.Getenv("GEOAPIFY_API_KEY")

	// Encode the query parameters
	query := url.Values{}
	query.Set("categories", category) // e.g., "catering.restaurant"
	query.Set("filter", filter)       // e.g., "circle:-87.770231,41.878968,5000"
	query.Set("apiKey", apiKey)

	// Create the request URL
	requestURL := fmt.Sprintf("https://api.geoapify.com/v2/places?%s", query.Encode())

	// Make the HTTP request
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Features []struct {
			Properties models.Place `json:"properties"`
		} `json:"features"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	// Extract the places from the result
	var places []models.Place
	for _, feature := range result.Features {
		places = append(places, feature.Properties)
	}

	return places, nil
}
