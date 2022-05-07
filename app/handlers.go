package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type House struct {
	ConstructionType     string `json:"construction_type"`
	NumberBedrooms       int    `json:"number_bedrooms"`
	NumberBathrooms      int    `json:"number_bathrooms"`
	Area                 int    `json:"area"`
	ConstructionYear     int    `json:"construction_year"`
	Condition            string `json:"condition"`
	EnergeticCertificate string `json:"energetic_certificate"`
	GarageSpots          bool   `json:"garage_spots"`
	ParkingLots          bool   `json:"parking_lots"`
	SwimmingPools        bool   `json:"swimming_pools"`
	Elevators            bool   `json:"elevators"`
	FullyEquipped        bool   `json:"fully_equipped"`
	Address              string `json:"address"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!!")
}

func getAllHouses(w http.ResponseWriter, r *http.Request) {
	houses := []House{
		{"apartament", 3, 3, 90, 2022, "new", "a", true, false, false, true, true, "Address 1"},
		{"apartament", 2, 1, 85, 2015, "second hand", "b", true, false, false, false, true, "Address 2"},
		{"apartament", 3, 2, 90, 2021, "second hand", "a", true, false, false, true, false, "Address 3"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}
