package figuras

type Circulo struct {
	radio float64
}

func (c *Circulo) Construct(radio float64) {
	c.radio = radio
}

const Pi float64 = 3.1416

func (c Circulo) Area() float64 {
	return Pi * c.radio * c.radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * Pi * c.radio
}
