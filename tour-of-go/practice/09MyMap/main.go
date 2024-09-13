package main

import (
	"fmt"
)

func main() {

	languages := make(map[string]int)

	languages["js"] = 1980
	languages["GO"] = 2007
	languages["ruby"] = 1990
	languages["c++"] = 1984

	fmt.Println(languages)
	fmt.Println("js:", languages["js"])
	delete(languages, "ruby")
	fmt.Println(languages)

	//loop in golang
	for _, value := range languages {
		fmt.Println(value)
	}

}
