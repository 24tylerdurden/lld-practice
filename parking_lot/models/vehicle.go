package models

type VehicleType string

const (
	Car   VehicleType = "Car"
	Truck VehicleType = "Truck"
	Bike  VehicleType = "Bike"
)

type Vehicle struct {
	Type         VehicleType
	LicensePlate string
}
