package figuras

// Struct

type Circulo struct{
	Radio float64
	Pi float64
}

// Metodos

func (c *Circulo) CalcularArea() float64 {
	area := c.Pi * (c.Radio * c.Radio)
	return area
}

func (c *Circulo) CalcularPerimetro() float64 {
	perimetro := 2*c.Pi * c.Radio
	return perimetro
}

