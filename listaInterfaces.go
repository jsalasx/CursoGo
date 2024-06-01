package main

import "fmt"

type figuras2D interface {
	area() float64
	perimetro() float64
}
type cuadrado struct {
	lado float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.lado * c.lado
}
func (c cuadrado) perimetro() float64 {
	return c.lado * 4
}

func (c cuadrado) String() string {
	return fmt.Sprintf("lado: %f", c.lado)
}

func (r rectangulo) String() string {
	return fmt.Sprintf("base: %f, altura: %f", r.base, r.altura)
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}
func (r rectangulo) perimetro() float64 {
	return r.base + r.base + r.altura + r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func calcularPerimetro(f figuras2D) {
	fmt.Println("Perimetro: ", f.perimetro())
}

func mainInrefacesLista() {
	myCuadrado := cuadrado{lado: 2}
	//fmt.Println(myCuadrado)
	//fmt.Println(myCuadrado.area())

	miRectangulo := rectangulo{base: 2, altura: 4}
	//fmt.Println(miRectangulo)
	//fmt.Println(miRectangulo.area())
	calcular(myCuadrado)
	calcular(miRectangulo)

	calcularPerimetro(myCuadrado)
	calcularPerimetro(miRectangulo)

	myInterface := []interface{}{"Hola", 12, 4.90}

	fmt.Println(myInterface...)

}
