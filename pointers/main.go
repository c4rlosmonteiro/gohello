package main

import "fmt"

func main() {
	myString := "My string!"
	fmt.Printf("My string value is '%s' and my pointer is '%p'\n", myString, &myString) // My string value is 'My string!' and my pointer is '0xc000096020'

	//here I have two different strings, the value was copied, the memory position of both is different
	myString2 := myString
	fmt.Printf("My string value is '%s' and my pointer is '%p'\n", myString2, &myString2)

	
}
