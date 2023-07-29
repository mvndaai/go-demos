package rectangle_test

import (
	"fmt"
	"testing"

	"github.com/mvndaai/go-interface-demo/reuseable/shape"
	"github.com/mvndaai/go-interface-demo/reuseable/shape/rectangle"
	"github.com/stretchr/testify/assert"
)

var _ shape.EqualDistant = rectangle.Rectangle{}

func TestEqualDistant(t *testing.T) {
	tests := []struct {
		lenght           float64
		expectedPermiter float64
		expectedArea     float64
	}{
		{
			lenght:           10,
			expectedPermiter: 40,
			expectedArea:     100,
		},
	}

	shape := rectangle.Rectangle{}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.lenght), func(t *testing.T) {
			actualArea := shape.Area(tt.lenght)
			assert.Equal(t, tt.expectedArea, actualArea)

			actualPerimeter := shape.Perimeter(tt.lenght)
			assert.Equal(t, tt.expectedPermiter, actualPerimeter)
		})
	}
}
