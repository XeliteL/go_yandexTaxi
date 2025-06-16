package main

type TripParameters struct {
	Distance float64
	Duration float64
}

const (
	pricePerKm     = 10.0
	pricePerMinute = 2.0
)

func CalculateBasePrice(t TripParameters) float64 {
	return t.Distance*pricePerKm + t.Duration*pricePerMinute
}
