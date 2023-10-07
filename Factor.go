package main

import (
	"fmt"
	"math"
)

// Shape is an interface for different types of shapes.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle is a concrete implementation of the Shape interface.
type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle is a concrete implementation of the Shape interface.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// ShapeFactory is a factory for creating different types of shapes.
type ShapeFactory struct{}

// CreateShape is a factory method that creates a Shape based on the given shapeType.
func (sf *ShapeFactory) CreateShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return &Circle{Radius: 5}
	case "rectangle":
		return &Rectangle{Width: 4, Height: 3}
	default:
		return nil
	}
}

func main() {
	factory := &ShapeFactory{}

	// Create a circle using the factory.
	circle := factory.CreateShape("circle")
	fmt.Printf("Circle - Area: %.2f, Perimeter: %.2f\n", circle.Area(), circle.Perimeter())

	// Create a rectangle using the factory.
	rectangle := factory.CreateShape("rectangle")
	fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", rectangle.Area(), rectangle.Perimeter())
}
