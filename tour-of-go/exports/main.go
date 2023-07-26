package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/exports/authops"
	"github.com/liberty234/lessons/tour-of-go/exports/buffers"
	"github.com/liberty234/lessons/tour-of-go/exports/data"
	"github.com/liberty234/lessons/tour-of-go/exports/datatime"
	"github.com/liberty234/lessons/tour-of-go/exports/deck"
	"github.com/liberty234/lessons/tour-of-go/exports/email"
	"github.com/liberty234/lessons/tour-of-go/exports/fileops"
	math "github.com/liberty234/lessons/tour-of-go/exports/math"
	"github.com/liberty234/lessons/tour-of-go/exports/network"
	"github.com/liberty234/lessons/tour-of-go/exports/parsers"
	"github.com/liberty234/lessons/tour-of-go/exports/parsing"
	"github.com/liberty234/lessons/tour-of-go/exports/somaphore"
	"github.com/liberty234/lessons/tour-of-go/exports/sorting"
	string1 "github.com/liberty234/lessons/tour-of-go/exports/string"
	"github.com/liberty234/lessons/tour-of-go/exports/test"
	"github.com/liberty234/lessons/tour-of-go/exports/util"
	"github.com/liberty234/lessons/tour-of-go/exports/xmlops"
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

	//used util as exported package
	bik := util.CarsAmount{Toyota: 9000, Lexus: 10000, Benx: 45000}
	fmt.Println("CarsTotal:", bik)
	util.Cars()

	//used network as exported package
	fmt.Println(network.Get)
	network.Network()

	//used sorting as exported package
	sorting.Sorting()

	//used xmlops as exported package
	fmt.Println(xmlops.Mlx)

	//used authops as exported package
	authops.Premiar()

	//used email as exported package
	email.Address()
	fmt.Println(email.ReceiverAddress)

	// used datatime as exported package
	datatime.Calculate()
	fmt.Println("the Datatime is", datatime.Datatime)

	//used somaphore as exported package
	somaphore.Ben()
	somaphore.Bigi()

}
