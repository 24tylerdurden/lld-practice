package models

import (
	"sync"
)

type Floor struct {
	mu            sync.Mutex
	FloorNumber   int
	Spots         []*ParkingSpot
	TotalSpots    int
	OccupiedSpots int
}

func NewFloor(number int, spots []*ParkingSpot) *Floor {
	return &Floor{
		FloorNumber: number,
		Spots:       spots,
		TotalSpots:  len(spots),
	}
}

func (f *Floor) FindAvailableSpot(vType VehicleType) *ParkingSpot {
	f.mu.Lock()
	defer f.mu.Unlock()

	for _, spot := range f.Spots {
		if spot.IsAvailable() && f.isSpotCompatible(spot.SpotType, vType) {
			return spot
		}
	}

	return nil
}

func (f *Floor) AssignVehicle(spotId string, vehicle *Vehicle) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	for _, spot := range f.Spots {
		if spot.ID == spotId {
			if spot.IsAvailable() {
				spot.AssignVehicle(vehicle)
				f.OccupiedSpots++
				return true
			}
		}
	}
	return false
}

func (f *Floor) RemoveVehicle(spotId string) bool {
	for _, spot := range f.Spots {
		if spot.ID == spotId && !spot.IsAvailable() {
			spot.RemoveVehicle()
			return true
		}
	}
	return false
}

func (f *Floor) isSpotCompatible(spotType SpotType, vType VehicleType) bool {
	switch vType {
	case Bike:
		return spotType == MotorCycle || spotType == Compact
	case Car:
		return spotType == Compact || spotType == Large || spotType == Electric
	case Truck:
		return spotType == Large
	default:
		return false
	}
}

func (f *Floor) GetOccupancyRate() float64 {
	if f.TotalSpots == 0 {
		return 0
	}
	return float64(f.OccupiedSpots) / float64(f.TotalSpots)
}
