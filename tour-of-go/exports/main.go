package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/exports/buffers"
	"github.com/liberty234/lessons/tour-of-go/exports/config"
	"github.com/liberty234/lessons/tour-of-go/exports/data"
	"github.com/liberty234/lessons/tour-of-go/exports/deck"
	"github.com/liberty234/lessons/tour-of-go/exports/fileops"
	math "github.com/liberty234/lessons/tour-of-go/exports/math"
	"github.com/liberty234/lessons/tour-of-go/exports/network"
	"github.com/liberty234/lessons/tour-of-go/exports/parsers"
	"github.com/liberty234/lessons/tour-of-go/exports/parsing"
	"github.com/liberty234/lessons/tour-of-go/exports/sorting"
	string1 "github.com/liberty234/lessons/tour-of-go/exports/string"
	"github.com/liberty234/lessons/tour-of-go/exports/test"
	"github.com/liberty234/lessons/tour-of-go/exports/util"
)

func main() {
	//used math export package
	fmt.Println(math.Person)
	math.PersonName()
	math.Add()

	//used data export package
	fmt.Println("Your Name:", data.Happy)
	data.FirstName()
	data.Class()

	//use test as export package
	test.Square()
	fmt.Println("the first person is", test.FirstPerson)

	//use string export package
	fmt.Println("my name is ", string1.Name)
	string1.Info()

	//use fileops export package
	fmt.Println("bill paid", fileops.Paybill)
	fileops.Bill()

	//using parsers as exported package
	parsers.Goods()
	parsers.Total()

	//using deck as exported package
	deck.Cards()
	fmt.Println("MYName:", deck.MyName)

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

	//used config as exported package
	fmt.Println("Game name:", config.GameName)
	config.Namevex()
	config.Calcu()

}
