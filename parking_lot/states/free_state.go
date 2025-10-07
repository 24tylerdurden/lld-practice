package states

import "LLD-PRACTICE/parking_lot/models"

type FreeState struct {
}

func (f *FreeState) AssignVehicle(spot *models.ParkingSpot, vehicle *models.Vehicle) {
	spot.Vehicle = vehicle
	spot.State = &OccupiedState{}
}

func (f *FreeState) RemoveVehicle(spot *models.ParkingSpot) {
	panic("Can't perform remove operation on empty state")
}

func (f *FreeState) IsAvailable() bool {
	return false
}
