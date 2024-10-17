package services

import (
	"github.com/frankschweitzer/Apartment-Scorer/pkg/models"
)

func FetchNearbyPlaces(address string) ([]models.Place, error) {
	// client, err := maps.NewClient(maps.WithAPIKey("YOUR_GOOGLE_API_KEY"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Logic to search nearby places based on address
	places := []models.Place{
		{
			Name:    "Gym XYZ",
			Address: "123 Fitness St",
			Rating:  4.5,
		},
	}
	return places, nil
}
