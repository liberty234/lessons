package main

import "fmt"

func GetPizza() {

	var pizza = map[string]string{
		"John": "pepperoni",
		"Mary": "Cheese",
		"Tim":  "Mushroom",
	}

	fmt.Println(pizza)
	fmt.Println(pizza["John"])
	pizza["John"] = "Peppers"
	fmt.Println(pizza)

	delete(pizza, "Tim")
	fmt.Println(pizza)

}
