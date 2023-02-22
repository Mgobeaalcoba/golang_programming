package main

import "fmt"

// Struct persona - ¿Como crearla?
type Persona struct{
	nombre string
	edad int
}

// Metodos - ¿Como crearlos? Ej: Metodo para la struct persona
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) editNombre() {
	fmt.Print("Ingrese su nuevo nombre => ")
	fmt.Scanln(&p.nombre)
}

// Herencia: En Golang se trabaja de una forma distinta a como se hace en Java (Kotlin tmb) y Python

type Empleado struct {
	// Tan simple como citar dentro la struct de la que hereda la nueva struct
	Persona
	Legajo int
}

func main() {
	// ¿Como crear objetos de tipo persona? 
	p1 := Persona{"Mariano",35}
	p2 := Persona{
		edad: 29,
		nombre: "Nicole",
	} // Al crear el objeto con sus keys entonces no necesito respetar el orden. 

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.nombre)
	fmt.Println(p2.nombre)

	p1.nombre = "Daniel"
	p2.nombre = "Antonella"

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.nombre)
	fmt.Println(p2.nombre)
	
	// Ejecuto el metodo de mi struct Persona creado arriba
	p1.imprimir()
	p2.imprimir()

	// Ejecuto el segundo metodo que me pide nombre por consola para cambiarlo

	//p1.editNombre()
	//p2.editNombre()
	//p1.imprimir()
	//p2.imprimir()

	// Voy a construir un nuevo objeto empleado que es un struct vacio por el momento: 

	em1 := Empleado{
		// Al tratarse de una herencia de Persona no puedo definir sus atributos desde la construcción del objeto Empleado
		//nombre: "Guillermo",
		//edad: 60,
		Legajo: 123456,
	}
	em1.nombre = "Guillermo"
	em1.edad = 60
	em1.Legajo = 132546

	fmt.Println(em1) // Imprime {} antes de la herencia, e imprime {{Guillermo 60}} luego de la herencia. Tiene dos pares de llaves. Con un atributo propio imprime: {{Guillermo 60} 132546} por lo que aquellos atributos que están dentro de doble par de llaves son los que trae de herencia. En este caso de Persona. s

	em1.imprimir()

}