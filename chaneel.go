package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func main() {

	c := make(chan string, 2)

	fmt.Println("hello")

	go say("world", c)

	fmt.Println(<-c)
}
