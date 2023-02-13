package main

import "fmt"

func main() {
	fmt.Println("Ingrese el valor de venta del producto:")
	var n1 float64
	fmt.Scanln(&n1)

	var igv float64 = impuestos(n1)
	fmt.Println("El IGV es de", igv)
	fmt.Println("El precio de venta es de", n1+igv)

}

func impuestos(num1 float64) float64 {
	return num1 * 0.18
}
