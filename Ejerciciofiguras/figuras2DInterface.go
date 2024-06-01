package figuras

type Figura2D interface {
	Area() float64
	Perimetro() float64
}

func CalcularArea(f Figura2D) float64 {
	return f.Area()
}
