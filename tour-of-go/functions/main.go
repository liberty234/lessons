package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/functions/auth"
	"github.com/liberty234/lessons/tour-of-go/functions/authetication"
	"github.com/liberty234/lessons/tour-of-go/functions/bac"
	"github.com/liberty234/lessons/tour-of-go/functions/calculate"
	"github.com/liberty234/lessons/tour-of-go/functions/colours"
	"github.com/liberty234/lessons/tour-of-go/functions/compressions"
)

func main() {
	FirstName, LastName := auth.GetFullName("Liberty", "Ebikade")
	fmt.Println(FirstName, LastName)

	Name := authetication.GetPersonName("grace")
	fmt.Println(Name)

	fmt.Println()

	class1, class2, class3, class4 := bac.GetClass("Boilogy", "English", "Chemistry", "Geography")
	fmt.Println(class1)
	fmt.Println(class2)
	fmt.Println(class3)
	fmt.Println(class4)

	y := calculate.GetCal(9, 6)
	fmt.Println(y)

	fmt.Println()

	color1, color2, color3, color4, color5 := colours.GetColors("Blue", "Green", "White", "Red", "yellow")
	fmt.Println(color1)
	fmt.Println(color2)
	fmt.Println(color3)
	fmt.Println(color4)
	fmt.Println(color5)

	h := compressions.GetMaths(4, 5, 8, 9)
	fmt.Println(h)

}
