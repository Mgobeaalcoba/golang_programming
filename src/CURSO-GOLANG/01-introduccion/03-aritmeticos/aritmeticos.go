package main

import "fmt"

func main() {
	a := 20
	b := 10

	// comentario
	/*
		comentario de bloque
	*/

	// suma
	result := a + b
	fmt.Println("Suma", result)

	// resta - alcanza solo con = xq la variable la inicialice arriba
	result = a - b
	fmt.Println("Resta", result)

	// multiplicación
	result = a * b
	fmt.Println("Multi", result)

	// división
	result = a / b
	fmt.Println("División", result)

	// problema de los tipados estaticos. 

	c := 3
	d := 2

	div := c/d

	fmt.Println("Div", div) // imprime 1 en lugar de 1,5 porque div es por default un int. Si quiero que impima 1,5 debo declararlo como un float y declarar dos nuevas variables de tipo float o transformar las que ya tengo. 

	e := 3.1
	f := 2.0

	var div2 float64 = e / f

	fmt.Println("Div", div2) 

	// modulo o sobrante de una división

	result = a % b
	fmt.Println("Modulo", result)

}