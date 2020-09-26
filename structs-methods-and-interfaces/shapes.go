package shape

import "math"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle Shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns a area of reactangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle Struct
type Circle struct {
	Radius float64
}

// Area returns a area of circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Triangle Shape
type Triangle struct {
	Base   float64
	Height float64
}

// Returns area of Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter returns perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// Area returns area of a rectangle
func Area(rectangle Rectangle) float64 {

	return rectangle.Height * rectangle.Width
}
