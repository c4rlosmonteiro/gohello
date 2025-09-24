package main

import (
	"fmt"
	"math"
)

func info() {
	//const

	const inflationRate1 = 2.5
	//or
	const inflationRate2 float64 = 2

	//var

}

func main() {
	const inflationRate = 6.5
	//This varibale will be assigned its default "null value" automatically
	//by Go. In case of a float64 value, tha's "0.0"
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5

	//To use scan you need to pass a pointer
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
