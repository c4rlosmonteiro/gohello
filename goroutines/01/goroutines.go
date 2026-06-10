package main

import (
	"fmt"
)

func main() {

	myIntegerChannel := make(chan int16)

	go func() {
		myIntegerChannel <- 1
		myIntegerChannel <- 2
		myIntegerChannel <- 3
	}()

	//The main funcion will wait until it receives a value from the channel or until the channel is closed.
	value := <- myIntegerChannel	
	fmt.Println("=> ", value)

	value2 := <- myIntegerChannel	
	fmt.Println("=> ", value2)

	fmt.Println("exit")
}