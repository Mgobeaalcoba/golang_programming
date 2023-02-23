package main

import (
	"fmt"
	"paquetes/figuras"
)

func main() {

	// Genero un cuadrado y un circulo

	cuadrado1 := figuras.Cuadrado{
		Ancho:  10.5,
		Altura: 15.3,
	}

	circulo1 := figuras.Circulo{
		Pi:    3.14,
		Radio: 9.7,
	}

	fmt.Printf("El area del cuadrado es de %.2f\n", figuras.CalculoDeArea(&cuadrado1))
	fmt.Printf("El perimetro del cuadrado es de %.2f\n", figuras.CalculoDePerimetro(&cuadrado1))
	fmt.Printf("El area del circulo es de %.2f\n", figuras.CalculoDeArea(&circulo1))
	fmt.Printf("El perimetro del circulo es de %.2f\n", figuras.CalculoDePerimetro(&circulo1))

}