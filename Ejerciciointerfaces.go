package main

import (
	"curso/figuras"
	"fmt"
)

func mainEjercicioInterfaces() {
	fmt.Println("Hello, World!")

	circulo := figuras.Circulo{}
	circulo.Construct(5.00)
	cuadrado := figuras.Cuadrado{Lado: 5}
	rectangulo := figuras.Rectangulo{Base: 5, Altura: 10}
	fmt.Println("Area del circulo: ", figuras.CalcularArea(circulo))

	fmt.Println("Area de cuadrado: ", figuras.CalcularArea(cuadrado))

	fmt.Println("Area del rectangulo:", figuras.CalcularArea(rectangulo))
}
