package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("please send a feedback:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Thanks for rating us %v star \n", numRating+1)
	}

	/*gen := "65"
	fmt.Println(gen)

	numconvert, err := strconv.ParseFloat(strings.TrimSpace(gen), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Great", numconvert)
	}*/
}
