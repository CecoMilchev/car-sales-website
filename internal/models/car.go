package models

import "time"

type Car struct {
	ID                int       `json:"id"`
	Make              string    `json:"make"`
	Model             string    `json:"model"`
	Power             uint      `json:"power"`
	Price             uint      `json:"price"`
	Mileage           uint      `json:"mileage"`
	EmissionsCategory uint8     `json:"emissionsCategory"`
	FirstRegistration time.Time `json:"firstRegistration"`
	Color             string    `json:"color"`
	CategoryID        int       `json:"categoryID"`
	FuelTypeID        int       `json:"fuelTypeID"`
	GearboxID         int       `json:"gearboxID"`
	//OwnerId int    `json:"ownerId"`
}
