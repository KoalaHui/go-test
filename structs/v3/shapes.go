package structs

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {

	return (r.width * r.height)
}
func (c Circle) Area() float64 {

	return math.Pi * c.Radius * c.Radius
}
