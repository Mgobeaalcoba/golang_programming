// mapas en golang es similar a diccionarios en python o a los objetos en javascript. No tienen orden y se organizan por clave y valor

package main

import "fmt"

func main() {
	// ¿Como inicializar una mapa?
	dias := make(map[int]string)
	// Dentro del make va a el tipo map. Dentro del [] el tipo de dato de la llave y fuera del [] el tipo de dato del valor

	fmt.Println(dias)
	// map[] --> Map vacio

	// ¿Como cargo datos en mi map? 
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miércoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"

	fmt.Println(dias)
	// map[1:Domingo 2:Lunes 3:Martes 4:Miércoles 5:Jueves 6:Viernes 7:Sabado] --> Map con contenido

	// ¿Como modificamos los datos de mi mapa?

	dias[7] = "SABADO"

	fmt.Println(dias)
	// map[1:Domingo 2:Lunes 3:Martes 4:Miércoles 5:Jueves 6:Viernes 7:SABADO] --> Map modificado

	// ¿Como eliminamos objetos de mi mapa? 

	delete(dias, 1)

	fmt.Println(dias)
	// map[2:Lunes 3:Martes 4:Miércoles 5:Jueves 6:Viernes 7:SABADO] --> Mapa con la key 1 eliminada.

	// Podemos consultar la longitud o len de una mapa pero no su capacidad ya que la misma es indefinida

	fmt.Println(dias, len(dias))
	// map[2:Lunes 3:Martes 4:Miércoles 5:Jueves 6:Viernes 7:SABADO] 6 --> Mapa con su len al lado

	// Mapas complejas: o mapas con arrays de valor

	estudiantes := make(map[string][]int)
	// Mapa con key de tipo de string y value de tipo slicen de enteros

	// ¿Como relleno a esta map "compleja"?
	estudiantes["Alex"] = []int{13,15,16}
	estudiantes["Mariano"] = []int{17,9,11}
	estudiantes["Nicole"] = []int{19,15,18}

	fmt.Println(estudiantes, len(estudiantes))
	// map[Alex:[13 15 16] Mariano:[17 9 11] Nicole:[19 15 18]] 3 --> Map con slicen de values y su largo. 

	// ¿Como consultar todos los values de una misma key? 
	fmt.Println(estudiantes["Alex"])
	// [13 15 16] --> Todos los values de Alex

	// ¿Y si solo quiero su primera nota? 
	fmt.Println(estudiantes["Alex"][0])
	// 13 --> Solo la primera nota de Alex

}