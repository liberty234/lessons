package main

import "fmt"

func main() {
	NoCount := 20

	var result string

	if NoCount < 20 {
		result = "less than"
	} else if NoCount > 20 {
		result = "greater "
	} else {
		result = "equal"
	}
	fmt.Println(result)

	if 9%3 == 0 {
		fmt.Println("the No is even")
	} else {
		fmt.Println("the No is odd")
	}

	if No := 3; No < 10 {
		fmt.Println("No is less than 10")
	} else {
		fmt.Println("No is greater than 10")
	}
}
