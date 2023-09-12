package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/functions/auth"
)

func main() {
	FirstName, LastName := auth.GetFullName("Liberty", "Ebikade")
	fmt.Println(FirstName, LastName)

}
