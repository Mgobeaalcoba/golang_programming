package figuras

// Interfaces

type Geometrico interface {
	CalcularArea() float64
	CalcularPerimetro() float64
}

// Función global para mi interface:

func CalculoDeArea(geometrico Geometrico) float64 {
	area := geometrico.CalcularArea()
	return area
}

func CalculoDePerimetro(geometrico Geometrico) float64 {
	perimetro := geometrico.CalcularPerimetro()
	return perimetro
}