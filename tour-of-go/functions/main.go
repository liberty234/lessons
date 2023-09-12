package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/functions/auth"
	"github.com/liberty234/lessons/tour-of-go/functions/authetication"
)

func main() {
	FirstName, LastName := auth.GetFullName("Liberty", "Ebikade")
	fmt.Println(FirstName, LastName)

	Name := authetication.GetPersonName("grace")
	fmt.Println(Name)

}
