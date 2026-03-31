package main

import "fmt"
import "math"

func sumTwoNumbers(a int, b int) int {
	return a + b
}

func sumTwoNumbers2(a int, b int) (sum int) {
	sum = a + b
	return
}

func sumTwoNumberWithDifference(a int, b int) (int, int) {
	diff := math.Abs(float64(a - b))

	return a + b, int(diff)
}

func main() {
	fmt.Println("Main!")

	fmt.Printf("Result of sumTwoNumbers: %d \n", sumTwoNumbers(1, 2))
	fmt.Printf("Result of sumTwoNumbers2: %d \n", sumTwoNumbers2(2, 3))
	result, diff := sumTwoNumberWithDifference(5, 3)
	fmt.Printf("Result of sumTwoNumbersWithDifference: %d and %d \n", result, diff)

	sumToNumbers3 := func (a int, b int) int {
		return a + b
	}

	fmt.Printf("Resutl of sumTowNumbers3: %d \n", sumToNumbers3(9, 1))
}