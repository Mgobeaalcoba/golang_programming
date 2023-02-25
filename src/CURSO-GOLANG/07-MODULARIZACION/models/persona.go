package models

// Struct publica con atributos privados:
type Persona struct {
	nombre string
	edad   int
}

// Constructor: Parecido a un metodo normal, debe ser publico:
func (p *Persona) Constructor(nombre string, edad int) {
	p.nombre = nombre
	p.edad = edad
}

// Getter y Setter
func (p *Persona) GetNombre() string {
	return p.nombre
}

func (p *Persona) GetEdad() int {
	return p.edad
}

func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Persona) SetEdad(edad int) {
	p.edad = edad
}
