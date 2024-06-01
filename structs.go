package main

import "fmt"

type car struct {
	brand string
	year  int
}

func mainStrucs() {

	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra manera
	var otroCarro car
	otroCarro.brand = "Ferrari"

	fmt.Println(otroCarro)

}
