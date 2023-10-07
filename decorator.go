package main

import "fmt"

type Coffee interface {
	Cost() float64
	Ingredients() string
}

type Espresso struct{}

func (e *Espresso) Ingredients() string {
	return "Espresso"
}

func (e *Espresso) Cost() float64 {
	return 100.0
}

type CoffeeDecorator struct {
	Coffee Coffee
}

func (cd *CoffeeDecorator) Ingredients() string {
	return cd.Coffee.Ingredients()
}

func (cd *CoffeeDecorator) Cost() float64 {
	return cd.Coffee.Cost()
}

type Milk struct {
	Decorator *CoffeeDecorator
}

func (m *Milk) Ingredients() string {
	return m.Decorator.Ingredients() + ", Milk"
}

func (m *Milk) Cost() float64 {
	return m.Decorator.Cost() + 20.0
}

type Whip struct {
	Decorator *CoffeeDecorator
}

func (w *Whip) Ingredients() string {
	return w.Decorator.Ingredients() + ", Whip"
}

func (w *Whip) Cost() float64 {
	return w.Decorator.Cost() + 30.0
}

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
