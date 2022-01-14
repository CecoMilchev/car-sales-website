package models

type Car struct {
	Id int `json:"id"`
	//OwnerId int    `json:"ownerId"`
	Make  string `json:"make"`
	Model string `json:"model"`
	// Category          string    `json:"category"`
	// Power             uint      `json:"power"`
	// Price             uint      `json:"price"`
	// Mileage           uint      `json:"mileage"`
	// FuelType          FuelType  `json:"fuelType"`
	// Gearbox           Gearbox   `json:"gearbox"`
	// EmissionsCategory uint8     `json:"emissionsCategory"`
	// FirstRegistration time.Time `json:"firstRegistration"`
	// Color             string    `json:"color"`
	// Description       string    `json:"description"`
	// Features          []string  `json:"features"`
}
