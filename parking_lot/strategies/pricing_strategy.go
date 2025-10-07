package strategies

import "time"

type PricingStrategy interface {
	CalculatePrice(entry, exit time.Time, vehicleType string) float64
}
