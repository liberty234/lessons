package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to us company"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Your feedback:")

	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for rating us %v star.\n", input)
	fmt.Printf("type of rating %T\n", input)

}
