package main

import "fmt"

type Details struct {
	Name       string
	Email      string
	Number     int
	HasGlasses bool
	Age        int
}

func main() {
	fmt.Println("Struct in Golang")

	details := Details{"Liberty", "liberty@mail.com", 7061267454, false, 18}
	fmt.Println(details)
	fmt.Printf("All details: %+v\n", details)

}
