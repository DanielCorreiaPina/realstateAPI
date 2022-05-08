package domain

import "github.com/DanielCorreiaPina/realstateAPI/errs"

type House struct {
	Id                   int    `json:"id"`
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

type HouseRepository interface {
	FindAll() ([]House, error)
	FindById(string) (*House, *errs.AppError)
}
