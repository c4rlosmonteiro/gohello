package main

import "fmt"

func main() {
	myStringChannel1 := make(chan string)
	myStringChannel2 := make(chan string)

	go func() {
		myStringChannel1 <- "my str channel 1"
	}()

	go func() {
		myStringChannel2 <- "my str channel 2"
	}()

	//select statement will block until one these cases can run
	select {
	case msgChannel1 := <- myStringChannel1:
		fmt.Println("msgChannel1:", msgChannel1)
	case msgChannel2 := <- myStringChannel2:
		fmt.Println("myChannel2:", msgChannel2)
	}
	
	fmt.Println("exit")
}
