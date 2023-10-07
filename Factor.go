package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

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

type ShapeFactory struct{}

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

	circle := factory.CreateShape("circle")
	fmt.Printf("Circle - Area: %.2f, Perimeter: %.2f\n", circle.Area(), circle.Perimeter())

	rectangle := factory.CreateShape("rectangle")
	fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", rectangle.Area(), rectangle.Perimeter())
}
