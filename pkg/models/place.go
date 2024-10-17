package models

type Place struct {
	Name    string  `json:"name"`
	Address string  `json:"formatted"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}
