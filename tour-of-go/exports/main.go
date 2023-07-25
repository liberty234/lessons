package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/exports/buffers"
	"github.com/liberty234/lessons/tour-of-go/exports/data"
	"github.com/liberty234/lessons/tour-of-go/exports/deck"
	"github.com/liberty234/lessons/tour-of-go/exports/fileops"
	math "github.com/liberty234/lessons/tour-of-go/exports/math"
	"github.com/liberty234/lessons/tour-of-go/exports/parsers"
	"github.com/liberty234/lessons/tour-of-go/exports/parsing"
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
	fmt.Println("the first person is", test.FirstPerson)

	//use string export package
	fmt.Println("my name is ", string1.Name)
	string1.Info()

	//use fileops export package
	don := fileops.Paybill{Fees: "100k", ElBill: "5k", Gas: "12k"}
	fmt.Println("Total paybill", don)
	fileops.Bill()

	//using parsers as exported package
	parsers.Goods()
	parsers.Total()

	//using deck as exported package
	deck.Cards()

	//used buffer as export package
	buffers.Register()

	//used parsing as export package
	parsing.Parsing()
	parsing.Divide()

}
