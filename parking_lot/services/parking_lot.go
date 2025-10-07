package services

import (
	"LLD-PRACTICE/parking_lot/models"
	"LLD-PRACTICE/parking_lot/states"
	"LLD-PRACTICE/parking_lot/strategies"
	"errors"
	"fmt"
	"sync"
	"time"
)

type ParkingLot struct {
	mu      sync.Mutex
	Floors  []*models.Floor
	Pricing strategies.PricingStrategy
}

var instance *ParkingLot
var Once sync.Once

func GetInstance() *ParkingLot {
	Once.Do(func() {
		instance = &ParkingLot{
			Pricing: &strategies.HourlyPricing{
				Rates: map[models.VehicleType]float64{
					models.Car:   5.0,
					models.Truck: 10.0,
					models.Bike:  2.5,
				},
			},
		}
		instance.initializeFloors(3, 50)
	})
	return instance
}

func (pl *ParkingLot) initializeFloors(numFloors, spotsPerFloor int) {
	pl.Floors = make([]*models.Floor, numFloors)

	for floorNumber := 0; floorNumber < numFloors; floorNumber++ {
		spots := make([]*models.ParkingSpot, spotsPerFloor)

		for i := 0; i < len(spots); i++ {
			spotType := models.Compact

			if i%10 == 0 {
				spotType = models.Large
			}

			if floorNumber == 2 && i&5 == 0 {
				spotType = models.Electric
			}

			spots[i] = &models.ParkingSpot{
				ID:          fmt.Sprintf("F%d-S%03d", floorNumber+1, i+1),
				SpotType:    spotType,
				FloorNumber: floorNumber + 1,
				State:       &states.FreeState{},
			}
		}
		pl.Floors[floorNumber] = models.NewFloor(floorNumber+1, spots)
	}

}

func (p *ParkingLot) ParkVehicle(vehicle *models.Vehicle) (*models.Ticket, error) {
	p.mu.Lock()

	defer p.mu.Unlock()

	for _, floor := range p.Floors {
		spot := floor.FindAvailableSpot(vehicle.Type)
		if spot != nil {
			if floor.AssignVehicle(spot.ID, vehicle) {
				return &models.Ticket{
					TicketId:     fmt.Sprintf("Ticket-%d", time.Now().UnixNano()),
					LicensePlate: vehicle.LicensePlate,
					EntryTime:    time.Now(),
					VehicleType:  vehicle.Type,
				}, nil

			}
		}
	}

	return nil, errors.New("no availabe spots for parcking")
}

func (p *ParkingLot) UnParkVehicle(ticket *models.Ticket) (float64, error) {
	p.mu.Lock()

	defer p.mu.Unlock()

	for _, floor := range p.Floors {

		if floor.RemoveVehicle(ticket.SpotId) {
			price := p.Pricing.CalculatePrice(
				ticket.EntryTime,
				time.Now(),
				string(ticket.VehicleType),
			)
			return price, nil
		}

	}

	return 0, errors.New("spots not found")
}
