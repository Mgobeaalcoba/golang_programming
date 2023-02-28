package main

import (
	"fmt"
	"net/http"
	"time"
)

// recibe la url del servidor que es un string
func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor) // Al mirar la function Get vemos que devuelve dos cosas: la respuesta (que no la quiero ahora por eso el _) y el error en caso de fallar la respuesta
	if err != nil {
		//fmt.Println(servidor, "No está disponible")
		canal <- servidor + "No está disponible" //Asigno mensaje a un canal de la misma forma en la que se asignan variables en R
	} else {
		//fmt.Println(servidor, "Esta funcionando")
		canal <- servidor + "Está funcionando"
	}
}

func main() {
	// Voy a verificar el tiempo de respuesta para estas llamadas a servidores. 

	inicio := time.Now() // Tiempo de inicio

	canal := make(chan string) // Creo un canal de string porque estamos usando URL´s. El canal lo tengo que enviar a mi función "revisarServidor"

	// El canal nos permite tenér la comunicación entre las distintas operaciones concurrentes que go run time está ejecutando a la vez. 

	// Armo una lista de servidores a verificar si están funcionando
	servidores := []string{
		"https://oregoom.com/",
		"https://www.udemy.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://www.google.com/",
	}

	// Paso los servidores 1 a 1. Es decir, sin concurrencia
	/*
	for i := 0; i < len(servidores); i++ {
		revisarServidor(servidores[i])
	}
	*/

	// Paso los servidores todos juntos. Es decir, con CONCURRENCIA:
	// Para ello debo agregar el go run time con la palabra reservada "go" delante de la función invocada.  

	for _ , servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	// Ciclo for para obtener los mensajes que me está devolviendo el servidor vía los canales. 
	for i := 0; i <= len(servidores); i++ {
		fmt.Println(<-canal)
	}

	// Con concurrencia la devolución de la validación de si el servidor funciona correctamente no es en el orden en que se envian sino que es en función de cual respuesta llega primero hasta la ultima. 

	// El problema con esto es que no va a terminar la ejecución de las consultas antes del tiempo de ejecución, por lo que el programa se cierra sin recibir la respuesta sobre si los servidores funcionan o no.
	
	// Para comunicarse entre las distintas consultas (hilos) y saber que nos está respondiendo el servidor debemos crear "channels" o canales. 

	tiempoEjecución2 := time.Since(inicio) // Tiempo desde el inicio al final

	fmt.Println("Tiempo de ejecución:", tiempoEjecución2)

}