package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/exports/data"
	math "github.com/liberty234/lessons/tour-of-go/exports/math"
	string1 "github.com/liberty234/lessons/tour-of-go/exports/string"
	"github.com/liberty234/lessons/tour-of-go/exports/test"
)

func main() {
	//used math export package
	fmt.Println(math.Person)
	math.PersonName()
	math.Add()

	//used data export package
	info := data.PersonalInfo{Name: "Liberty", Phnumber: "07061267494", DateOfBirth: 4, Age: 22}
	fmt.Println(info)
	data.Class()

	//use test as export package
	test.Square()

	//use string export package
	fmt.Println("my name is ", string1.Name)
	string1.Info()

}