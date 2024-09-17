package main

import "fmt"

func GetAges() {

	ages := make(map[string]int)

	ages["james"] = 25
	fmt.Printf("James is %v years old\n", ages["james"])

	grade := map[string]float32{
		"james": 3.6,
		"susan": 3.9,
	}
	fmt.Printf("James GPA is %2f\n", grade["james"])
	fmt.Printf("Susan GPA is %2f\n", grade["susan"])

	var Visited map[string]bool
	Visited = make(map[string]bool)
	Visited["A"] = true
	fmt.Printf("A has been visited? %v\n", Visited["A"])

}
