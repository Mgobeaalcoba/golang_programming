// mensajes es un paquete que tiene un archivo que se llama saludar.go

package mensajes // Indico a que paquete corresponde el codigo
import "fmt"

// Las funciones de los paquetes deben llevar su inicial en mayus
// Si la inicial de la función es en minuscula significa que es una función privada y no puede ser accedida desde otro archivo. Por ejemplo el main.

func Hola() {
	fmt.Println("Hola desde el paquete mensaje")
}

// Ahora desde nuestro archivo main vamos a importar este paquete como hemos realizado con "fmt" siempre

// Prueba de modificadores de acceso: 

const mensaje = "Hola desde mi const mensaje privada" // Al igual que con las funciones si yo declaro mi variable o constante con minus es privada y no puede ser accedida por fuera del archivo. Pero si usada. Si la declaro en mayuscula puede ser accedida y usada. 

func Imprimir() {
	fmt.Println(mensaje)
	funcionPrivada()
}

func funcionPrivada() {
	fmt.Println("Hola desde la función privada!!!")
}

const Mensaje2 = "Hola desde mi const Mensaje2 publica"

