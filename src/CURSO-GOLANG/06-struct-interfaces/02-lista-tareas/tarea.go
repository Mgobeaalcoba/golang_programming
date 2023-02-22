package main

import "fmt"

// Lista de tareas: Slicen en el que vamos a guardar objetos de tipo tarea
type TaskList struct{
	tasks [] *Task
}

// Creo entonces metodos para agregar y metodos para borrar tareas
func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) deleteTask(index int) {
	// Remover una tarea de mi lista: ¿Como? Generando una nueva lista que se forma con dos fragmentos: desde el inicio hasta el indice enviado y desde el indice+1 hasta el final
	tl.tasks = append(tl.tasks[:index],tl.tasks[index+1:]...)
}

// Armo mi struct Task
type Task struct {
	name      string
	desc      string
	completed bool
}

// Armo los metodos de mi struct Task
func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripción: %s\nCompletado: %t\n", t.name, t.desc, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}

func main() {

	// Genero las tareas

	t1 := Task{
		name: "Curso de Go",
		desc: "Completar curso de Go este mes",
		completed: false,
	}

	t2 := Task{
		name: "Curso de HTML",
		desc: "Completar curso de HTML esta semana",
		completed: true,
	}

	t1.toPrint()
	t2.toPrint()

	// Genero las lista de tareas

	lista := TaskList{}
	lista.appendTask(&t1) // Mando la referencia y no una copia con el &
	lista.appendTask(&t2) // Mando la referencia y no una copia con el &

	fmt.Println(lista)
	// Imprime los espacios en memoria

	t3 := Task{
		name: "Curso de CSS",
		desc: "Completar curso de CSS esta semana",
		completed: true,
	}

	lista.appendTask(&t3)

	fmt.Println(lista)
	// Sigue imprimiendo los espacios en memoria. ¿Como hacer para que imprima cada uno de los objetos almacenados? Con un ciclo for

	for index, task := range lista.tasks{
		fmt.Println(index, task) // Imprime el indice y el contenido del objeto
		task.toPrint() // Usa el metodo toPrint para cada task
	}

	/*
	Veamos ahora: 

	0 &{Curso de Go Completar curso de Go este mes false}
	Nombre: Curso de Go
	Descripción: Completar curso de Go este mes
	Completado: false
	1 &{Curso de HTML Completar curso de HTML esta semana true}
	Nombre: Curso de HTML
	Descripción: Completar curso de HTML esta semana
	Completado: true
	2 &{Curso de CSS Completar curso de CSS esta semana true}
	Nombre: Curso de CSS
	Descripción: Completar curso de CSS esta semana
	Completado: true

	*/

	// Si solo quiero acceder a una entonces: 

	fmt.Println("Imprimirlas de a una: ")
	lista.tasks[2].toPrint()

	/*

	Imprimirlas de a una:
	Nombre: Curso de CSS
	Descripción: Completar curso de CSS esta semana
	Completado: true

	*/

	// Eliminar una tarea: 

	lista.deleteTask(1)

	for _, task := range lista.tasks{
		task.toPrint() // Usa el metodo toPrint para cada task
	}

}