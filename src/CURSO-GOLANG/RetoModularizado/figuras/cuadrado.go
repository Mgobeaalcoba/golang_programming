package figuras

// Struct

type Cuadrado struct{
	Ancho float64
	Altura float64
}

// Metodos

func (c *Cuadrado) CalcularArea() float64 {
	area := c.Ancho * c.Altura
	return area
}

func (c *Cuadrado) CalcularPerimetro() float64 {
	perimetro := 2*c.Ancho * 2*c.Altura
	return perimetro
}
