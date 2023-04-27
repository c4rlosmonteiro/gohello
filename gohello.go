package main

import "fmt"

func generateHelloMessage(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	message := generateHelloMessage("world")
	fmt.Println(message)
}
