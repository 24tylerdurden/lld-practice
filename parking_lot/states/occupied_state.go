package states

import "LLD-PRACTICE/parking_lot/models"

type OccupiedState struct{}

func (o *OccupiedState) AssignVehicle(spot *models.ParkingSpot, vehicle *models.Vehicle) {
	panic("can't assign vehicle to occupied spot")
}

func (o *OccupiedState) RemoveVehicle(spot *models.ParkingSpot) {
	spot.Vehicle = nil
	spot.State = &FreeState{}
}

func (o *OccupiedState) IsAvailable() bool {
	return false
}
