package main

import "fmt"

// Coffee interface defines the methods for cost and ingredients.
type Coffee interface {
	Cost() float64
	Ingredients() string
}

// Espresso is a concrete implementation of Coffee.
type Espresso struct{}

func (e *Espresso) Ingredients() string {
	return "Espresso"
}

func (e *Espresso) Cost() float64 {
	return 100.0
}

// CoffeeDecorator is a base decorator struct for Coffee.
type CoffeeDecorator struct {
	Coffee Coffee
}

func (cd *CoffeeDecorator) Ingredients() string {
	return cd.Coffee.Ingredients()
}

func (cd *CoffeeDecorator) Cost() float64 {
	return cd.Coffee.Cost()
}

// Milk is a concrete decorator for adding milk to Coffee.
type Milk struct {
	Decorator *CoffeeDecorator
}

func (m *Milk) Ingredients() string {
	return m.Decorator.Ingredients() + ", Milk"
}

func (m *Milk) Cost() float64 {
	return m.Decorator.Cost() + 20.0
}

// Whip is a concrete decorator for adding whip to Coffee.
type Whip struct {
	Decorator *CoffeeDecorator
}

func (w *Whip) Ingredients() string {
	return w.Decorator.Ingredients() + ", Whip"
}

func (w *Whip) Cost() float64 {
	return w.Decorator.Cost() + 30.0
}

// Chocolate is a concrete decorator for adding chocolate to Coffee.
type Chocolate struct {
	Decorator *CoffeeDecorator
}

func (c *Chocolate) Ingredients() string {
	return c.Decorator.Ingredients() + ", Chocolate"
}

func (c *Chocolate) Cost() float64 {
	return c.Decorator.Cost() + 50.0
}

func main() {
	espresso := &Espresso{}
	cappuccino := &Milk{Decorator: &CoffeeDecorator{Coffee: &Whip{Decorator: &CoffeeDecorator{Coffee: espresso}}}}
	cappuccinoWithChocolate := &Chocolate{Decorator: &CoffeeDecorator{Coffee: &Milk{Decorator: &CoffeeDecorator{Coffee: espresso}}}}

	fmt.Println(espresso.Ingredients())
	fmt.Println(espresso.Cost())

	fmt.Println(cappuccino.Ingredients())
	fmt.Println(cappuccino.Cost())

	fmt.Println(cappuccinoWithChocolate.Ingredients())
	fmt.Println(cappuccinoWithChocolate.Cost())
}
