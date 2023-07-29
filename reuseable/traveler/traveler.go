//go:generate mockgen -destination=./mocks/traveler.go -package=mocks -source=./traveler.go

package traveler

import "time"

type Traveler interface {
	Travel(time.Duration) float64
}
