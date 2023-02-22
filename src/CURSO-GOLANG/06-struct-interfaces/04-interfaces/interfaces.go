/*
interfaces: Sirve para manejar los metodos identicos de las estructuras
*/

package main

import "fmt"

// Genero mi interface que agrupa a mis estructuras:

type Animal interface {
	mover() string // El interface Animal va a manejar el metodo mover que retorna un string.
}

// Genero mis estructuras

type Perro struct{

}

type Pez struct{

}

type Pajaro struct{

}

// A cada una de estas struct le quiero agregar el metodo "mover" ya que todos pueden hacerlo

func (*Perro) mover() string {
	return "Soy perro y camino"
}

func (*Pez) mover() string {
	return "Soy pez y nado"
}

func (*Pajaro) mover() string {
	return "Soy pajaro y vuelo"
}

// Ahora voy a generar una función global para manejar mi interface: 

func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}

func main() {

	perro := Perro{}
	moverAnimal(&perro)

	pez := Pez{}
	moverAnimal(&pez)

	pajaro := Pajaro{}
	moverAnimal(&pajaro)

	/*
	Soy perro y camino
	Soy pez y nado
	Soy pajaro y vuelo
	*/
}