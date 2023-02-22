package main

import "fmt"

// Creamos dos estructuras para establecer una relación de 1:1 entre ellas

type User struct{
	nombre string
	email string
	activo bool
}

type Student struct{
	// En la herencia yo declaro como atributo directamente a la estructura que deseo invocar. En la relación creo una variable a la que le paso como type el nombre de la estructura con la que lo quiero relacionar
	user User // Un estudiante solo puede tener 1 usuario y 1 usuario solo puede tener 1 estudiante. 
	codigo string
}

// Creamos dos structs (curso y videos) para establecer una relación 1:muchos entre curso y videos

type Curso struct{
	titulo string
	videos []Video // Con un slicen del struct Video declaro relación 1:muchos para curso.
}

type Video struct{
	titulo string
	curso Curso // Acá declaro que la relación de un video es con solo 1 Curso
}

func main() {

	/* Relacion uno a uno: 
	// Creo dos users. Y luego voy a hacer que solo uno sea estudiante: 

	alex := User{
		nombre: "Alex",
		email: "alex@gmail.com",
		activo: true,
	}

	mariano := User{
		nombre: "Mariano",
		email: "mariano@gmail.com",
		activo: false,
	}

	// Hago que solo alex sea estudiante

	alexStudent := Student{
		user: alex, // Debo pasar como valor de user, un objeto User
		codigo: "001asf5",
	}

	fmt.Println(alex)
	fmt.Println(mariano)
	fmt.Println(alexStudent)

	
	{Alex alex@gmail.com true}
	{Mariano mariano@gmail.com false}
	{{Alex alex@gmail.com true} 001asf5} // Se imprime igual que con la herencia, pero tiene una diferencia fundamental a la hora de trabajar con los atributos. Si era herencia podía acceder al mail, dato del user, desde alexStudent.email pero al estar relacionado debo acceder así: 
	

	fmt.Println()
	fmt.Println(alexStudent.user.email)
	*/

	/* Relacion uno a muchos:
	
	*/

	video1 := Video{
		titulo: "01-Introducción",
	}
	video2 := Video{titulo: "02-Instalación"}

	curso1 := Curso{
		titulo: "Curso sobre Golang de Cero a Master",
		videos: []Video{video1,video2},
	}

	// Para declarar el atributo curso de mis videos lo debo hacer debajo de la inicializacion de mi curso. Sino arrojará error ya que se compila de arriba a abajo el programa

	video1.curso = curso1
	video2.curso = curso1

	fmt.Println(video1)
	fmt.Println(video2)
	fmt.Println(curso1)

	// ¿Como accedo a cada titulo de los video del curso1? 

	fmt.Println()

	for _, video := range curso1.videos {
		fmt.Println(video.titulo)
	}
}	
