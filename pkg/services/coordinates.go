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

func FetchLatLonFromAddress(address string) (float64, float64, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	apiKey := os.Getenv("GEOAPIFY_API_KEY")

	// Create the request URL for geocoding
	geocodeURL := fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&apiKey=%s", url.QueryEscape(address), apiKey)

	// Make the request
	resp, err := http.Get(geocodeURL)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	// Parse the response
	var result models.GeocodeResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, 0, err
	}

	// Extract latitude and longitude
	if len(result.Features) > 0 {
		lat := result.Features[0].Properties.Lat
		lon := result.Features[0].Properties.Lon
		return lat, lon, nil
	}

	return 0, 0, fmt.Errorf("no geocoding results found")
}
