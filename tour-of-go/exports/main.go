package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/exports/auth"
	"github.com/liberty234/lessons/tour-of-go/exports/buffers"
	"github.com/liberty234/lessons/tour-of-go/exports/colors"
	"github.com/liberty234/lessons/tour-of-go/exports/config"
	"github.com/liberty234/lessons/tour-of-go/exports/csv"
	"github.com/liberty234/lessons/tour-of-go/exports/data"
	"github.com/liberty234/lessons/tour-of-go/exports/datetime"
	"github.com/liberty234/lessons/tour-of-go/exports/db"
	"github.com/liberty234/lessons/tour-of-go/exports/deck"
	"github.com/liberty234/lessons/tour-of-go/exports/encryption"
	"github.com/liberty234/lessons/tour-of-go/exports/fileops"
	"github.com/liberty234/lessons/tour-of-go/exports/logs"
	"github.com/liberty234/lessons/tour-of-go/exports/network"
	"github.com/liberty234/lessons/tour-of-go/exports/parsers"
	"github.com/liberty234/lessons/tour-of-go/exports/parsing"
	"github.com/liberty234/lessons/tour-of-go/exports/sorting"
	string1 "github.com/liberty234/lessons/tour-of-go/exports/string"
	"github.com/liberty234/lessons/tour-of-go/exports/subs"
	"github.com/liberty234/lessons/tour-of-go/exports/test"
	"github.com/liberty234/lessons/tour-of-go/exports/util"
	"github.com/liberty234/lessons/tour-of-go/exports/validation"
)

func main() {
	//used math export package
	fmt.Println(db.Person)
	db.PersonName()
	db.Add()

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
	fileops.Billpaid()

	//using parsers as exported package
	parsers.Goods()
	parsers.Total()

	//using deck as exported package
	fmt.Println("MYName:", deck.MyName)
	deck.Cards()
	deck.NewDeck()

	//used buffer as export package
	fmt.Println("Level:", buffers.Level)
	buffers.Register()
	buffers.Sub()

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
	network.NetworkName()

	//used sorting as exported package
	fmt.Println("Ans:", sorting.Sort)
	sorting.Sorting()
	sorting.Books()

	//used config as exported package
	fmt.Println("Game name:", config.GameName)
	config.Namevex()
	config.Calcu()

	//used auth as exported package
	fmt.Println("Multipo:", auth.Multi)
	auth.MyNaame()
	auth.Deva()

	//used colors as exported package
	fmt.Println("MyBestColor:", colors.BestColor)
	colors.Colors()
	colors.Colour()

	//used validation as exported package
	fmt.Println("Valid NO:", validation.Valid)
	validation.Values()
	validation.Grade()

	//used datetime as exported package
	fmt.Println("Time:", datetime.TimeDate)
	datetime.DateTime()
	datetime.Weeks()

	//used logs as exported package
	fmt.Println("Log in time:", logs.LogTime)
	logs.LogName()
	logs.UserName()

	//used subs as exported package
	fmt.Println("Sub Ans =", subs.Sub)
	subs.SubFunx()
	subs.Manus()

	//used csv as exported package
	fmt.Println("Yuor Name :", csv.YourName)
	csv.LastName()
	csv.MiddleName()

	//used encryption as exported package
	fmt.Println("Ans:", encryption.Tans)
	encryption.TansAns()
	encryption.Tanfun()

}
