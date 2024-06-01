package mypackage

import "fmt"

type Pcbre struct {
	ram   int
	disk  int
	brand string
}

func (myPC Pcbre) Ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *Pcbre) DuplicaterRAM() {
	myPC.ram = myPC.ram * 2
	myPC.disk = myPC.disk + 2
}

func (myPc *Pcbre) InitPc(ram int, disk int, brand string) {
	myPc.ram = ram
	myPc.disk = disk
	myPc.brand = brand
}
