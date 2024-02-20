package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greeter()

	//result := adder(65, 34)
	//fmt.Println("Result =", result)

	fmt.Println("result", adder(23, 84))
	fmt.Println("Result:", proAdder(23, 83, 93, 29))

}

func greeter() {
	fmt.Println("welcome to golang")
}

func adder(add1 int, add int) int {
	return add1 + add

}

func proAdder(add ...int) int {
	total := 0

	for _, ad := range add {
		total += ad
	}

	return total
}
