package main

import "fmt"

func main() {
	var strValues [3]string
	//the default value for string is empty, the below line of code will print 3 empty spaces
	fmt.Println("values of strValues:", strValues) // values of strValues: [  ]

	letters := [4]string {"a", "b", "c", "d"}
	fmt.Println("Letters array:", letters) // Letters array: [a b c d]

	slice := letters[1:2] // from position 1 to position 2 (2 value is not inclued)
	fmt.Println("Slice:", slice) // Slice: [b]

	letters2 := letters[:] // get all elements
	fmt.Println("Letters 2:", letters2)

	// slices are connected with their original array 

	letters3 := [4]string {"a", "b", "c", "d"}
	sliceFromLetter3 := letters3[1:2] 
	//letters3= [a b c d]   sliceFromLetter3= [b]
	fmt.Println("letters3=", letters3, " ", "sliceFromLetter3=", sliceFromLetter3)
	sliceFromLetter3[0] = "z"
	//letters3= [a z c d]   sliceFromLetter3= [z]
	fmt.Println("letters3=", letters3, " ", "sliceFromLetter3=", sliceFromLetter3)
}
