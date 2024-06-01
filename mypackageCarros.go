package mypackage

import "fmt"

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// Print message para imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}
