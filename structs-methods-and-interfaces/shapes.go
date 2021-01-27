package interfaces

import "math"

// Shape is a interface
type Shape interface {
	Area() float64
}

// Rectangle defines the width and height of the rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area return the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

//Circle defines the radius of a circle
type Circle struct {
	Radius float64
}

// Area return the area of a Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle defines the base and height of a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area return the area of a Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter return the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
