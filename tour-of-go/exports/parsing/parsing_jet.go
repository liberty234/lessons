package parsing

import "fmt"

var Par int = 65 + 12

func GetPassing() {
	Parsing := [...]string{"partletter", "partgrowth"}
	fmt.Println("Parsing package:", Parsing)

}

func GetDivide() {
	fmt.Println("Ans:", div)
	div := getDivide(80, 20)
	fmt.Println("Divition;", div)

}
