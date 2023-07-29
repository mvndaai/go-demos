package shape

type EquildistantArea interface {
	Area(l float64) float64
}

type EquildistantPerimeter interface {
	Perimeter(l float64) float64
}

type EqualDistant interface {
	EquildistantArea
	EquildistantPerimeter
}
