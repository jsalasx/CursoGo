package main

//awesome-go.com

import "fmt"

func message(text string, c chan string) {
	c <- text

}

func main() {

	c := make(chan string, 2)

	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println(len(c), cap(c))
	close(c)
	for message := range c {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)

	go message("Mensaje 1", email1)
	go message("Mensaje 2", email2)

	//defer close(c)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Recibido en email1: ", m1)
		case m2 := <-email2:
			fmt.Println("Recibido en email2: ", m2)
		}

	}
}
