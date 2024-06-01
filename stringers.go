package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPc pc) String() string {
	return fmt.Sprintf("PC: %s, RAM: %d, Disk: %d", myPc.brand, myPc.ram, myPc.disk)
}

func main() {

	myPc := pc{ram: 16, brand: "Lenovo", disk: 512}
	fmt.Println(myPc)

}
