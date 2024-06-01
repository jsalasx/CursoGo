package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println(text)
}

func mainConcurrencia() {

	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)

	go say("World", &wg)
	wg.Wait()

	go func(text string) {
		fmt.Println("Anonymous->" + text)
	}("papas")
	time.Sleep(time.Second * 1)

}
