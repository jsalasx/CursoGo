package main

import (
	"fmt"
)

func mainMaps () {
	m:= make(map[string]int)

	m["jose"] = 14
	m["pepito"] = 20

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	value, ok := m["jose"]
	fmt.Println(value, ok)
}