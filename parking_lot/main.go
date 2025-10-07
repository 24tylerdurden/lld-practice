package parkinglot

import (
	"LLD-PRACTICE/parking_lot/models"
	"LLD-PRACTICE/parking_lot/services"
	"fmt"
	"time"
)

func main() {
	lot := services.GetInstance()

	car := &models.Vehicle{
		LicensePlate: "ABC123",
		Type:         models.Car,
	}

	ticket, err := lot.ParkVehicle(car)

	if err != nil {
		fmt.Printf("Parking Failed: %v\n", err)
		return
	}

	fmt.Println("Parked %s in spot %s\n", car.LicensePlate, ticket.SpotId)

	// simulate the sleep

	time.Sleep(2 * time.Second)

	// Unpark

	price, err := lot.UnParkVehicle(ticket)
	if err != nil {
		fmt.Printf("Unparking failed: %v\n", err)
		return
	}

	fmt.Printf("Pay $%.2f for parking \n", price)
}
