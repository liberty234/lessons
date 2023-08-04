package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/exports/auth"
	"github.com/liberty234/lessons/tour-of-go/exports/authetication"
	"github.com/liberty234/lessons/tour-of-go/exports/booking"
	"github.com/liberty234/lessons/tour-of-go/exports/buffers"
	"github.com/liberty234/lessons/tour-of-go/exports/calulate"
	"github.com/liberty234/lessons/tour-of-go/exports/colors"
	"github.com/liberty234/lessons/tour-of-go/exports/comprssion"
	"github.com/liberty234/lessons/tour-of-go/exports/concurrent"
	"github.com/liberty234/lessons/tour-of-go/exports/config"
	"github.com/liberty234/lessons/tour-of-go/exports/csv"
	"github.com/liberty234/lessons/tour-of-go/exports/data"
	"github.com/liberty234/lessons/tour-of-go/exports/databases"
	"github.com/liberty234/lessons/tour-of-go/exports/datetime"
	"github.com/liberty234/lessons/tour-of-go/exports/db"
	"github.com/liberty234/lessons/tour-of-go/exports/deck"
	"github.com/liberty234/lessons/tour-of-go/exports/dp"
	"github.com/liberty234/lessons/tour-of-go/exports/email"
	"github.com/liberty234/lessons/tour-of-go/exports/encryption"
	"github.com/liberty234/lessons/tour-of-go/exports/fileops"
	"github.com/liberty234/lessons/tour-of-go/exports/files"
	"github.com/liberty234/lessons/tour-of-go/exports/helpers"
	"github.com/liberty234/lessons/tour-of-go/exports/http"
	"github.com/liberty234/lessons/tour-of-go/exports/imageops"
	"github.com/liberty234/lessons/tour-of-go/exports/jp"
	"github.com/liberty234/lessons/tour-of-go/exports/json"
	"github.com/liberty234/lessons/tour-of-go/exports/jsonops"
	"github.com/liberty234/lessons/tour-of-go/exports/logs"
	"github.com/liberty234/lessons/tour-of-go/exports/network"
	"github.com/liberty234/lessons/tour-of-go/exports/operation"
	"github.com/liberty234/lessons/tour-of-go/exports/parsers"
	"github.com/liberty234/lessons/tour-of-go/exports/parsing"
	"github.com/liberty234/lessons/tour-of-go/exports/regexops"
	"github.com/liberty234/lessons/tour-of-go/exports/saver"
	"github.com/liberty234/lessons/tour-of-go/exports/saverops"
	"github.com/liberty234/lessons/tour-of-go/exports/score"
	"github.com/liberty234/lessons/tour-of-go/exports/sort"
	"github.com/liberty234/lessons/tour-of-go/exports/sorting"
	"github.com/liberty234/lessons/tour-of-go/exports/state"
	"github.com/liberty234/lessons/tour-of-go/exports/statistics"
	string1 "github.com/liberty234/lessons/tour-of-go/exports/string"
	"github.com/liberty234/lessons/tour-of-go/exports/structure"
	"github.com/liberty234/lessons/tour-of-go/exports/subs"
	"github.com/liberty234/lessons/tour-of-go/exports/tags"
	"github.com/liberty234/lessons/tour-of-go/exports/test"
	"github.com/liberty234/lessons/tour-of-go/exports/timeops"
	"github.com/liberty234/lessons/tour-of-go/exports/util"
	"github.com/liberty234/lessons/tour-of-go/exports/validation"
	"github.com/liberty234/lessons/tour-of-go/exports/webs"
	"github.com/liberty234/lessons/tour-of-go/exports/xml"
	"github.com/liberty234/lessons/tour-of-go/exports/xmlops"
)

func main() {
	//used math export package
	fmt.Println(db.Person)
	db.GetPersonName()
	db.GetAdd()

	//used data export package
	fmt.Println("Your Name:", data.Happy)
	data.FetchFirstName()
	data.FetchClass()

	//use test as export package
	test.GetMainSquare()
	test.GetSquare()
	fmt.Println("the first person is", test.FirstPerson)

	//use string export package
	fmt.Println("my name is ", string1.Name)
	string1.GetPersonalInfo()
	string1.GetInfo()

	//use fileops export package
	fmt.Println("bill paid", fileops.PayBill)
	fileops.GetBill()
	fileops.GetBillPaid()

	//using parsers as exported package
	fmt.Println("Total:", parsers.Goo)
	parsers.GetGoods()
	parsers.GetTotal()

	//using deck as exported package
	fmt.Println("MYName:", deck.MyName)
	deck.GetCards()
	deck.FetchNewDeck()

	//used buffer as export package
	fmt.Println("Level:", buffers.Level)
	buffers.GetRegister()
	buffers.GetSub()

	//used parsing as export package
	fmt.Println("Ans", parsing.Par)
	parsing.GetPassing()
	parsing.GetDivide()

	//used util as exported package
	fmt.Println("Car:", util.Cars)
	util.GetCarAmount()
	util.GetCar()

	//used network as exported package
	fmt.Println(network.Get)
	network.FetchNetwork()
	network.FetchNetworkName()

	//used sorting as exported package
	fmt.Println("Ans:", sorting.Sort)
	sorting.GetSorting()
	sorting.GetBooks()

	//used config as exported package
	fmt.Println("Game name:", config.GameName)
	config.FetchNameVet()
	config.FetchCalcu()

	//used auth as exported package
	fmt.Println("Multipo:", auth.Multi)
	auth.GetMyNaame()
	auth.GetDeva()

	//used colors as exported package
	fmt.Println("MyBestColor:", colors.BestColor)
	colors.FetchColors()
	colors.FetchColour()

	//used validation as exported package
	fmt.Println("Valid NO:", validation.Valid)
	validation.GetValues()
	validation.GetGrade()

	//used datetime as exported package
	fmt.Println("Time:", datetime.TimeDate)
	datetime.GetDateTime()
	datetime.GetWeeks()

	//used logs as exported package
	fmt.Println("Log in time:", logs.LogTime)
	logs.GetLogName()
	logs.GetUserName()

	//used subs as exported package
	fmt.Println("Sub Ans =", subs.Sub)
	subs.GetSubFunx()
	subs.GetManus()

	//used csv as exported package
	fmt.Println("Yuor Name :", csv.YourName)
	csv.GetLastName()
	csv.GetMiddleName()

	//used encryption as exported package
	fmt.Println("Ans:", encryption.Tans)
	encryption.FetchTansAns()
	encryption.FetchTanFun()

	//used score as exported package
	fmt.Println("MainScore:", score.MainScore)
	score.GetSource()
	score.GetTotalScore()

	//used regexops as exported package
	fmt.Println("Gate:", regexops.Build)
	regexops.GetHouse()
	regexops.GetCal()

	//used files as exported package
	fmt.Println(files.Files)
	files.GetNameOfFiles()
	files.GetValuesCheck()

	//used statistics as exported package
	fmt.Println("Ans:", statistics.Static)
	statistics.GetValue()
	statistics.GetStaticvalues()

	//used state as exported package
	fmt.Println("State of origin:", state.StateOfOrigin)
	state.GetStateName()
	state.GetOrigin()

	//used compression as exported package
	fmt.Println("Camp:", comprssion.Comp)
	comprssion.GetCampMaths()
	comprssion.GetXmlMaths()

	//used xml as exported package
	fmt.Println(xml.Mathx)
	xml.GetXmlPackage()
	xml.GetDoc()

	//used timeops as exported package
	fmt.Println("levo:", timeops.Level)
	timeops.GetCarNumber()
	timeops.GetUserName()

	//used operation as exported package
	fmt.Println("Oprami", operation.Opramin)
	operation.GetOpera()
	operation.GetDivide()

	//used tags as exported package
	fmt.Println("TagVar:", tags.Tags)
	tags.GetTagFun()
	tags.GetTagsTypeFunc()

	//used booking as exported package
	fmt.Println("booked:", booking.Booking)
	booking.FetchBooked()
	booking.GetTicketBooked()

	//used email as exported package
	fmt.Println("\nEnter email:", email.Email)
	email.GetEmailFunc()
	email.GetValidEmail()

	//used json as exported package
	fmt.Println(json.Opramin)
	json.GetOpera()
	json.GetDivide()

	//used json as exported package
	fmt.Println("Occurent:", concurrent.Occurent)
	concurrent.GetNumber()
	concurrent.GetManus()

	//used authetication as exported package
	fmt.Println("SUrName:", authetication.FulllName)
	authetication.GetPersonName()
	authetication.FetchAdd()

	//used webs as exported package

	fmt.Println("NewClient:", webs.NewName)
	webs.GetPotName()
	webs.GetValues()

	//used saver as exported package
	fmt.Println("Saver;", saver.Saved)
	saver.GetWebName()
	saver.GetWebsite()

	//used  http as exported package
	fmt.Println("Http::", http.Game)
	http.GetView()
	http.GetWeb()

	//used databases as exported package
	fmt.Println(databases.Data)
	databases.GetDataBase()
	databases.GetBase()

	//used calulate as exported package
	fmt.Println("Ans:", calulate.Cal)
	calulate.GetCall()
	calulate.GetValue()

	//used imageops as exported package
	fmt.Println(imageops.Image)
	imageops.FetchImageops()
	imageops.FetchView()

	//used jsonops as exported package
	fmt.Println(jsonops.Name)
	jsonops.GetFunc()
	jsonops.GetJp()

	//used jp as exported package
	fmt.Println("Total:", jp.Jps)
	jp.GetJp()
	jp.GetNus()

	//used structure as exported package
	fmt.Println("Ans:", structure.Struct)
	structure.FetchGb()
	structure.FetchSolve()

	//used sort as exported package
	fmt.Println("Name:", sort.Name)
	sort.GetNameVex()
	sort.GetForm()

	//used helpers as exported package
	fmt.Println("Cost:", helpers.Total)
	helpers.GetCost()
	helpers.GetTotals()

	//used dp as exported package
	fmt.Println("UserName:", dp.UserName)
	dp.GetFulllName()
	dp.GetSet()

	//used xmlops as exported package
	fmt.Println("Xpos:", xmlops.Xpos)
	xmlops.GetPackage()
	xmlops.GetSort()

	//used svaerops as exported package
	fmt.Println(saverops.NameFull)
	saverops.GetFirstName()
	saverops.GetClass()

}
