package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}

type circulo struct {
	radious float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radious * c.radious
}

type retangulo struct {
	largura, altura float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func main() {
	c := circulo{
		radious: 10,
	}
	r := retangulo{largura: 10, altura: 10}

	fmt.Println("Area do circulo ->", c.area())
	fmt.Println("Area do retangulo ->", r.area())
}
