package triangle

import "math"

type Triangle struct{}

func (t Triangle) Area(l float64) float64 {
	area := (math.Sqrt(3) / 4) * math.Pow(l, 2)
	return math.Round(area*100) / 100
}

func (t Triangle) Perimeter(l float64) float64 {
	return l * 3
}
