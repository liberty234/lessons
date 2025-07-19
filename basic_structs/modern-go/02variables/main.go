package main

import "fmt"

const Hieght float64 = 75.5

func main() {
	// string with var
	var name string = "liberty"
	fmt.Println("my name is", name)
	fmt.Printf("the data type is %T\n", name)

	// writing var with boolean
	var isname bool = true
	fmt.Println("my name is", isname)
	fmt.Printf("the data type is %T\n", isname)

	// defining a float with var
	var p float64 = 34.45
	var j float64 = 7.5
	n := p / j
	fmt.Println("result:", n)

	fmt.Println(Hieght)
	fmt.Printf("the data type is %T\n", Hieght)

}
