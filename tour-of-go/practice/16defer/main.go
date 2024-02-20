package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	myDefer()

}

func myDefer() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}
