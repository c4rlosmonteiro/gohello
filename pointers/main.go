package main

import "fmt"

func main() {
	myString := "My string!"
	fmt.Printf("My string value is '%s' and my pointer is '%p'\n", myString, &myString) //out: My string value is 'My string!' and my pointer is '0xc000096020'

	//here I have two different strings, the value was copied, the memory position of both is different
	myString2 := myString
	fmt.Printf("My string value is '%s' and my pointer is '%p'\n", myString2, &myString2)

	fmt.Println("____________")

	myFloatVar := 5.4
	fmt.Printf("myFloatVar value is: %.2f | %p\n", myFloatVar, &myFloatVar) //out: myFloatVar value is: 5.40 | 0xc0000ac010
	var myFloatVarPointer *float64 = &myFloatVar
	// we can use the symbol '*' to assign a new value to the pointer location
	*myFloatVarPointer = 7.8
	fmt.Printf("myFloatVarPointer value is: %.2f | %p\n", *myFloatVarPointer, myFloatVarPointer) //out: myFloatVarPointer value is: 7.80 | 0xc0000ac010

	fmt.Println("____________")
}
