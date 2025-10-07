package models

import "time"

type Ticket struct {
	TicketId     string
	LicensePlate string
	EntryTime    time.Time
	SpotId       string
	VehicleType  VehicleType
}
