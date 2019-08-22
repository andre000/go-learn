package shapes

import "math"

// Shape Generic element with Area() function
type Shape interface {
	Area() float64
}

// Rectangle Struct with Width and Height
type Rectangle struct {
	Width  float64
	Height float64
}

// Area return the area of a Rectangle
func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}

// Circle Struct with Radius
type Circle struct {
	Radius float64
}

// Area return the area of a Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle Struct with Base and Height
type Triangle struct {
	Base   float64
	Height float64
}

// Area return the area of a Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter returns the perimeter of a Rectangle
func Perimeter(ret Rectangle) float64 {
	return (ret.Width + ret.Height) * 2
}

// Area return the area of a Rectangle
func Area(ret Rectangle) float64 {
	return (ret.Height * ret.Width)
}
