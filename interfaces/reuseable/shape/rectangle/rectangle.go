package rectangle

import "math"

type Rectangle struct{}

func (r Rectangle) Area(l float64) float64 {
	area := l * l
	return math.Round(area*100) / 100
}

func (r Rectangle) Perimeter(l float64) float64 {
	return l * 4
}
