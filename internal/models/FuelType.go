package models

type FuelType uint8

const (
	Petrol FuelType = iota
	Diesel
	HybridDiesel
	HybridPetrol
	Electric
	LPG
)

func (ft FuelType) String() string {
	switch ft {
	case HybridDiesel:
		return "Hybrid (diesel/electric)"
	case HybridPetrol:
		return "Hybrid (petrol/electric)"
	}

	return []string{"Petrol", "Diesel", "Electric", "LPG"}[ft]
}
