package models

type SpotState interface {
	AssignVehicle(spot *ParkingSpot, vehicle *Vehicle)
	RemoveVehicle(spot *ParkingSpot)
	IsAvailable() bool
}
