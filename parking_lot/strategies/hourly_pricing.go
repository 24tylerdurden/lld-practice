package strategies

import (
	"LLD-PRACTICE/parking_lot/models"
	"time"
)

type HourlyPricing struct {
	Rates map[models.VehicleType]float64
}

func (h *HourlyPricing) CalculatePrice(entry, exit time.Time, vehicleType string) float64 {
	duration := exit.Sub(entry)

	hours := duration.Hours()

	if hours < 1 {
		hours = 1
	}

	return hours * h.Rates[models.VehicleType(vehicleType)]
}
