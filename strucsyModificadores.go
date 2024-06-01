package main

import (
	"curso/mypackage"
	"fmt"
)

func mainStrucsYModificadores() {

	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	mypackage.PrintMessage("hola platzi")
}
