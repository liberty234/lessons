package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/exports/auth"
	"github.com/liberty234/lessons/tour-of-go/exports/booking"
	"github.com/liberty234/lessons/tour-of-go/exports/buffers"
	"github.com/liberty234/lessons/tour-of-go/exports/colors"
	"github.com/liberty234/lessons/tour-of-go/exports/comprssion"
	"github.com/liberty234/lessons/tour-of-go/exports/concurrent"
	"github.com/liberty234/lessons/tour-of-go/exports/config"
	"github.com/liberty234/lessons/tour-of-go/exports/csv"
	"github.com/liberty234/lessons/tour-of-go/exports/data"
	"github.com/liberty234/lessons/tour-of-go/exports/datetime"
	"github.com/liberty234/lessons/tour-of-go/exports/db"
	"github.com/liberty234/lessons/tour-of-go/exports/deck"
	"github.com/liberty234/lessons/tour-of-go/exports/email"
	"github.com/liberty234/lessons/tour-of-go/exports/encryption"
	"github.com/liberty234/lessons/tour-of-go/exports/fileops"
	"github.com/liberty234/lessons/tour-of-go/exports/files"
	"github.com/liberty234/lessons/tour-of-go/exports/json"
	"github.com/liberty234/lessons/tour-of-go/exports/logs"
	"github.com/liberty234/lessons/tour-of-go/exports/network"
	"github.com/liberty234/lessons/tour-of-go/exports/operation"
	"github.com/liberty234/lessons/tour-of-go/exports/parsers"
	"github.com/liberty234/lessons/tour-of-go/exports/parsing"
	"github.com/liberty234/lessons/tour-of-go/exports/regexops"
	"github.com/liberty234/lessons/tour-of-go/exports/score"
	"github.com/liberty234/lessons/tour-of-go/exports/sorting"
	"github.com/liberty234/lessons/tour-of-go/exports/state"
	"github.com/liberty234/lessons/tour-of-go/exports/statistics"
	string1 "github.com/liberty234/lessons/tour-of-go/exports/string"
	"github.com/liberty234/lessons/tour-of-go/exports/subs"
	"github.com/liberty234/lessons/tour-of-go/exports/tags"
	"github.com/liberty234/lessons/tour-of-go/exports/test"
	"github.com/liberty234/lessons/tour-of-go/exports/timeops"
	"github.com/liberty234/lessons/tour-of-go/exports/util"
	"github.com/liberty234/lessons/tour-of-go/exports/validation"
	"github.com/liberty234/lessons/tour-of-go/exports/xml"
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
	test.MainSquare()
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

	//used score as exported package
	fmt.Println("MainScore:", score.MainScore)
	score.Source()
	score.TotalScore()

	//used regexops as exported package
	fmt.Println("Gate:", regexops.Build)
	regexops.House()
	regexops.Cal()

	//used files as exported package
	fmt.Println(files.Files)
	files.NameOfFiles()
	files.ValuesCheck()

	//used statistics as exported package
	fmt.Println("Ans:", statistics.Static)
	statistics.Value()
	statistics.Staticvalues()

	//used state as exported package
	fmt.Println("State of origin:", state.StateOfOrigin)
	state.StateName()
	state.Origin()

	//used compression as exported package
	fmt.Println("Camp:", comprssion.Comp)
	comprssion.CampMaths()
	comprssion.Xmlmaths()

	//used xml as exported package
	fmt.Println(xml.Mathx)
	xml.Xmlpackage()
	xml.Doc()

	//used timeops as exported package
	fmt.Println("levo:", timeops.Level)
	timeops.CarNumber()
	timeops.UserName()

	//used operation as exported package
	fmt.Println("Oprami", operation.Opramin)
	operation.Opera()
	operation.Divide()

	//used tags as exported package
	fmt.Println("TagVar:", tags.Tags)
	tags.Tagfun()
	tags.TagsTypeFunc()

	//used booking as exported package
	fmt.Println("booked:", booking.Booking)
	booking.Booked()
	booking.TicketBooked()

	//used email as exported package
	fmt.Println("\nEnter email:", email.Email)
	email.Emailfunc()
	email.ValidEmail()

	//used json as exported package
	fmt.Println(json.Opramin)
	json.Opera()
	json.Divide()

	//used json as exported package
	fmt.Println("Occurent:", concurrent.Occurent)
	concurrent.Number()
	concurrent.Manus()

}
