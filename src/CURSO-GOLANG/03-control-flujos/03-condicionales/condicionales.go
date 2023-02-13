package main

import "fmt"

func main() {
	/** App para restaurante
	* Descuentos de 10% hasta 100 de consumo
	* Descuentos de 20% hasta 200 de consumo
	* Descuentos de 30% mayor a 200 de consumo
	* igv 19%
	 */
	var consumo, descuento, igv float64
	var datosDescuento string

	igv = 0.19

	// Entrada de datos
	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo > 0 {
		if consumo <= 100 {
			datosDescuento = "10%"
			descuento = 0.1
		} else if consumo <= 200 {
			descuento = 0.2
			datosDescuento = "20%"
		} else {
			descuento = 0.3
			datosDescuento = "30%"
		}
		fmt.Println("------- Factura de consumo -------")
		fmt.Printf("Su consumo es de %v,\nsu descuento es de %v,\ny su total a pagar sin impuestos es de %v,\nsu igv a pagar es de: %v,\nsu total final con impuestos es de %v", consumo, datosDescuento, consumo * (1 - descuento), (consumo * (1 - descuento)) * igv, (consumo * (1 - descuento)) + ((consumo * (1 - descuento)) * igv))
	} else {
		fmt.Println("Error al ingresar el consumo")
	}

}