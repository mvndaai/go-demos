package bike

import (
	"time"
)

type Manuel struct {
	AverageSpeed float64
}

func (b Manuel) Travel(d time.Duration) float64 {
	return b.AverageSpeed * float64(d)
}

type Electric struct {
	AverageSpeed float64
	BatteryLeft  time.Duration
}

func (e *Electric) Travel(d time.Duration) float64 {
	if e.BatteryLeft > d { // Drain a portion
		e.BatteryLeft -= d
	} else { // Drain everything and stop
		d = e.BatteryLeft
		e.BatteryLeft = 0
	}
	return e.AverageSpeed * float64(d)
}
