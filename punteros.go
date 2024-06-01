package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicaterRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {

	a := 50
	b := &a

	fmt.Println("b puntero:", b)
	fmt.Println("b valor:", *b)

	*b = 100

	fmt.Println("a:", a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	myPC.duplicaterRAM()
	fmt.Println(myPC)

	myPC.duplicaterRAM()
	fmt.Println(myPC)
}
