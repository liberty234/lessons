package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("swich and case in golang")

	rand.Seed(time.Now().UnixNano())
	DiceNumber := rand.Intn(6) + 1
	fmt.Println("Value:", DiceNumber)

	switch DiceNumber {
	case 1:
		fmt.Println("move to 1 spot")
	case 2:
		fmt.Println("move to 2 spot")
	case 3:
		fmt.Println("move to 3 spot")
	case 4:
		fmt.Println("move to 4 spot")
	case 5:
		fmt.Println("move to 5 spot")
	case 6:
		fmt.Println("you can open in 6 spot")
	}
}
