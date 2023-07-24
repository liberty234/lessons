package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/exports/data"
	math "github.com/liberty234/lessons/tour-of-go/exports/math"
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

}
