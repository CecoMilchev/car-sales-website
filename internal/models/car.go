package models

import (
	"time"
)

type Car struct {
	ID                int       `json:"id"`
	CategoryID        int       `json:"category_id"`
	FuelTypeID        int       `json:"fuelType_id" gorm:"column:fuelType_id"`
	GearboxID         int       `json:"gearbox_id"`
	Make              string    `json:"make"`
	Model             string    `json:"model"`
	Power             uint      `json:"power"`
	Mileage           uint      `json:"mileage"`
	EmissionsCategory uint8     `json:"emissionsCategory" gorm:"column:emissionsCategory"`
	FirstRegistration time.Time `json:"firstRegistration" gorm:"column:firstRegistration"`
	Color             string    `json:"color"`
	//OwnerId int    `json:"ownerId"`
}
