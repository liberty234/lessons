package main

import "fmt"

var logIn = 93092939

const LogInCode = "hs73snjsjk" //pulic

func main() {
	var surName string = "Ebikade"
	fmt.Println(surName)
	fmt.Printf("My surname is %T \n", surName)

	var isloogedIn bool = false
	fmt.Println(isloogedIn)
	fmt.Printf("variables is of type %T \n", isloogedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variables is of type %T \n", smallval)

	var smallfloat float64 = 255.233299393
	fmt.Println(smallfloat)
	fmt.Printf("variables is of type %T \n", smallfloat)

	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("variables is of type %T \n", anotherVar)

	fmt.Println(LogInCode)
	fmt.Printf("variables is of type %T \n", LogInCode)

	fmt.Println(logIn)
	fmt.Printf("variables is of type %T \n", logIn)

}
