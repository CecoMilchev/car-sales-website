package models

type Gearbox uint8

const (
	Manual Gearbox = iota
	Automatic
)

func (s Gearbox) String() string {
	switch s {
	case Manual:
		return "Manual Gearbox"
	case Automatic:
		return "Automatic Gearbox"
	}

	return "unknown"
}
