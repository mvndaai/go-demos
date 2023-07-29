package bike_test

import (
	"testing"
	"time"

	"github.com/mvndaai/go-interface-demo/reuseable/traveler"
	"github.com/mvndaai/go-interface-demo/reuseable/traveler/bike"
	"github.com/stretchr/testify/assert"
)

// These error if the types don't match the interface
var _ traveler.Traveler = &bike.Manuel{}
var _ traveler.Traveler = &bike.Electric{}

func TestManuelTravel(t *testing.T) {
	tests := []struct {
		name             string
		averageSpeed     float64
		tripTime         time.Duration
		expectedDistance float64
	}{
		{
			name:             "1 hour",
			averageSpeed:     10,
			tripTime:         time.Hour,
			expectedDistance: 3.6e+13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bike.Manuel{
				AverageSpeed: tt.averageSpeed,
			}
			actual := b.Travel(tt.tripTime)
			assert.Equal(t, tt.expectedDistance, actual)
		})
	}
}

func TestElectricTravel(t *testing.T) {
	tests := []struct {
		name            string
		averageSpeed    float64
		batteryDuration time.Duration
		tripTime        time.Duration

		expectedBattery  time.Duration
		expectedDistance float64
	}{
		{
			name:            "1 hour, enough batter",
			averageSpeed:    10,
			batteryDuration: 2 * time.Hour,
			tripTime:        time.Hour,

			expectedBattery:  time.Hour,
			expectedDistance: 3.6e+13,
		},
		{
			name:            "1 hour, lower batter",
			averageSpeed:    10,
			batteryDuration: time.Hour / 2,
			tripTime:        time.Hour,

			expectedBattery:  0,
			expectedDistance: 3.6e+13 / 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bike.Electric{
				AverageSpeed: tt.averageSpeed,
				BatteryLeft:  tt.batteryDuration,
			}
			actualDistance := b.Travel(tt.tripTime)
			assert.Equal(t, tt.expectedDistance, actualDistance)
			assert.Equal(t, tt.expectedBattery, b.BatteryLeft)
		})
	}
}
