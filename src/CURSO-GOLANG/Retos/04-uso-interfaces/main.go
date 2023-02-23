package main

import "fmt"

// Interfaces

type Geometrico interface {
	calcularArea() float64
	calcularPerimetro() float64
}

// Structs

type Cuadrado struct{
	ancho float64
	altura float64
}

type Circulo struct{
	radio float64
	pi float64
}

// Metodos: 

func (c *Cuadrado) calcularArea() float64 {
	area := c.ancho * c.altura
	return area
}

func (c *Cuadrado) calcularPerimetro() float64 {
	perimetro := 2*c.ancho * 2*c.altura
	return perimetro
}

func (c *Circulo) calcularArea() float64 {
	area := c.pi * (c.radio * c.radio)
	return area
}

func (c *Circulo) calcularPerimetro() float64 {
	perimetro := 2*c.pi * c.radio
	return perimetro
}

// Función global para mi interface:

func calculoDeArea(geometrico Geometrico) float64 {
	area := geometrico.calcularArea()
	return area
}

func calculoDePerimetro(geometrico Geometrico) float64 {
	perimetro := geometrico.calcularPerimetro()
	return perimetro
}


func main() {

	// Genero un cuadrado y un circulo

	cuadrado1 := Cuadrado{
		ancho: 10.5,
		altura: 15.3,
	}

	circulo1 := Circulo{
		pi: 3.14,
		radio: 9.7,
	}

	fmt.Printf("El area del cuadrado es de %.2f\n", calculoDeArea(&cuadrado1))
	fmt.Printf("El perimetro del cuadrado es de %.2f\n", calculoDePerimetro(&cuadrado1))
	fmt.Printf("El area del circulo es de %.2f\n", calculoDeArea(&circulo1))
	fmt.Printf("El perimetro del circulo es de %.2f\n", calculoDePerimetro(&circulo1))
}