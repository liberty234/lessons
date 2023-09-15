package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/functions/auth"
	"github.com/liberty234/lessons/tour-of-go/functions/authetication"
	"github.com/liberty234/lessons/tour-of-go/functions/bac"
	"github.com/liberty234/lessons/tour-of-go/functions/books"
	"github.com/liberty234/lessons/tour-of-go/functions/bufa"
	"github.com/liberty234/lessons/tour-of-go/functions/calculate"
	"github.com/liberty234/lessons/tour-of-go/functions/colours"
	"github.com/liberty234/lessons/tour-of-go/functions/compressions"
	"github.com/liberty234/lessons/tour-of-go/functions/concurents"
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

	dc := concurents.GetNumber(45, 8, 60)
	fmt.Println("Ans:", dc)

	fmt.Println()

	reg1, reg2, reg3, reg4 := bufa.GetRegister("Ebikade", "Liberty", "Best", "Peace")
	fmt.Println(reg1)
	fmt.Println(reg2)
	fmt.Println(reg3)
	fmt.Println(reg4)

	fmt.Println()

	booked1, booked2, booked3, booked4, booked5, booked6 := books.FetchBooked(15, 9, 6, 12, 25, 13)
	fmt.Println(booked1)
	fmt.Println(booked2)
	fmt.Println(booked3)
	fmt.Println(booked4)
	fmt.Println(booked5)
	fmt.Println(booked6)

}
