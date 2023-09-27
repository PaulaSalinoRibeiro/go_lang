package main

import (
	"fmt"
	"math"
)

type formas interface {
	area() float64
}

func printArea(f formas) {
	fmt.Printf("a area é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(2, c.raio)
}

func main() {

	r := retangulo{12, 10}
	printArea(r)

	c := circulo{3}
	printArea(c)

}
