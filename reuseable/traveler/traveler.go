package traveler

import "time"

type Traveler interface {
	Travel(time.Duration) float64
}
