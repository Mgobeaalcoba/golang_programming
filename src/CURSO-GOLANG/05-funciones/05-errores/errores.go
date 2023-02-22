// El manejo de errores, como se lo conoce en otros lenguajes como Python en Golang no existe. Por lo que debemos prepararlo de forma manual y nos podemos ayudar del package "errors" para ello.

package main

import (
	"errors"
	"fmt"
)

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir entre 0!!!")
	} else {
		return dividendo / divisor, nil // nil se usa para cuando queremos usar un valor ausente
	}
}

func main() {
	result, error := division(4, 0)

	if error == nil {
		fmt.Println("Resultado: ", result)
	} else {
		fmt.Println(error)
	}
}

/*
Si divido por cero me aparece el siguiente error:

panic: runtime error: integer divide by zero
*/
