package main

import "fmt"

func main() {

	// key   value
	menu := map[string]float64{
		"Espresso": 2.50,
		"Latte":    1.95,
	}

	fmt.Println("My map: ", menu)
	fmt.Println("Get specific element (Latte): ", menu["Latte"])
	// updating latte price
	menu["Latte"] = 3.95
	fmt.Println("New price (Latte): ", menu["Latte"])
	//adding new item
	menu["New coffee"] = 9.99
	//map len
	fmt.Println("The map len is ", len(menu), ", with the following values ", menu)

	//checking key existence 
	price, exists := menu["Cappuccino"]
	fmt.Println("The value associted with the key Cappuccino is ", price, ". This field existence is ", exists)

	//removing key
	delete(menu, "New coffee")
	fmt.Println("My map: ", menu)

}
