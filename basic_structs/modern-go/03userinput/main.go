package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("welcne to user input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate our Pizza:")

	//comma ok // err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us,", input)
}
