package main

import "fmt"

/*
	A closure in Go is a function value that references variables from its surrounding scope, 
	even after the outer function has finished executing. This allows functions to "remember" 
	and access or modify private data, effectively maintaining state between calls.
 */
func counter(i int) func () int {
	count := i

	return func () int {
		count += 1
		return count
	}
}

func main() {
	counter := counter(0)

	fmt.Println("Counter after first call: ", counter())
	fmt.Println("Counter after second call: ", counter())
	fmt.Println("Counter after third call: ", counter())
}
