package main

import (
	"fmt"
)

func main() {
	
	//aarray

	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	slice := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Apped
	slice = append(slice, 7)
	fmt.Println(slice)

	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
