package main

import "fmt"

func GetMenu() {

	menu := map[string]float64{
		"soup":  4.84,
		"pie":   7.82,
		"salad": 6.4,
		"Beans": 3.54,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])
	for k, v := range menu {
		fmt.Println(k, v)

	}

	phoneBook := map[int]string{
		+2347061267494: "liberty",
		+2340702874847: "Best",
		+2348947374774: "peace",
	}
	fmt.Println(phoneBook)
	for k, v := range phoneBook {
		fmt.Println(k, v)

	}
	phoneBook[+2347061267494] = "lee"
	fmt.Println(phoneBook)

}
