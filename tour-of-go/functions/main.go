package main

import (
	"fmt"

	"github.com/liberty234/lessons/tour-of-go/functions/auth"
	"github.com/liberty234/lessons/tour-of-go/functions/authetication"
	"github.com/liberty234/lessons/tour-of-go/functions/bac"
	"github.com/liberty234/lessons/tour-of-go/functions/bd"
	"github.com/liberty234/lessons/tour-of-go/functions/books"
	"github.com/liberty234/lessons/tour-of-go/functions/bufa"
	"github.com/liberty234/lessons/tour-of-go/functions/calculate"
	"github.com/liberty234/lessons/tour-of-go/functions/colours"
	"github.com/liberty234/lessons/tour-of-go/functions/compressions"
	"github.com/liberty234/lessons/tour-of-go/functions/concurents"
	"github.com/liberty234/lessons/tour-of-go/functions/configs"
	"github.com/liberty234/lessons/tour-of-go/functions/csv"
	"github.com/liberty234/lessons/tour-of-go/functions/database"
	"github.com/liberty234/lessons/tour-of-go/functions/datas"
	"github.com/liberty234/lessons/tour-of-go/functions/deck"
	"github.com/liberty234/lessons/tour-of-go/functions/dp"
	"github.com/liberty234/lessons/tour-of-go/functions/emails"
	"github.com/liberty234/lessons/tour-of-go/functions/encryptions"
	"github.com/liberty234/lessons/tour-of-go/functions/fileop"
	"github.com/liberty234/lessons/tour-of-go/functions/files"
	"github.com/liberty234/lessons/tour-of-go/functions/helper"
	"github.com/liberty234/lessons/tour-of-go/functions/http"
	"github.com/liberty234/lessons/tour-of-go/functions/imageop"
	"github.com/liberty234/lessons/tour-of-go/functions/jp"
	"github.com/liberty234/lessons/tour-of-go/functions/json"
	"github.com/liberty234/lessons/tour-of-go/functions/jsonop"
	"github.com/liberty234/lessons/tour-of-go/functions/logs"
	"github.com/liberty234/lessons/tour-of-go/functions/network"
	"github.com/liberty234/lessons/tour-of-go/functions/operation"
	"github.com/liberty234/lessons/tour-of-go/functions/parsers"
	"github.com/liberty234/lessons/tour-of-go/functions/parsing"
	"github.com/liberty234/lessons/tour-of-go/functions/regexop"
	"github.com/liberty234/lessons/tour-of-go/functions/saver"
	"github.com/liberty234/lessons/tour-of-go/functions/saverops"
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

	fmt.Println()

	game1, game2, game3, game4 := configs.FetchNameVet("Boxing", "Football", "volley Ball", "Table Tennis")
	fmt.Println(game1)
	fmt.Println(game2)
	fmt.Println(game3)
	fmt.Println(game4)

	name := bd.GetPersonName("Liberty Ebikade")
	fmt.Println(name)
	fmt.Println()
	FirstNames, LastNames, middleName := csv.GetName("Best", "Ebikade", "Ayemereh")
	fmt.Println(FirstNames)
	fmt.Println(LastNames)
	fmt.Println(middleName)

	d := database.GetCa(65, 902)
	fmt.Println(d)

	bv := datas.GetNum(4, 5, 8, 9)
	fmt.Println("Ans:", bv)
	fmt.Println()

	ace, five, four, two := deck.GetCard("Spades", "Diamond", "Spades", "Diamond")
	fmt.Println(ace)
	fmt.Println(five)
	fmt.Println(four)
	fmt.Println(two)
	fmt.Println()

	FirstName1, LastName1, middleName1 := dp.GetAllName("Godbless", "Okoh", "Obi")
	fmt.Println(FirstName1)
	fmt.Println(LastName1)
	fmt.Println(middleName1)
	fmt.Println()

	company1, company2, company3 := emails.FetchCompanyEmail("globalbestCompany@gmail.com", "GodBlessgroup@gmail.com", "libertyInfoCompany@gmail.com")
	fmt.Println(company1)
	fmt.Println(company2)
	fmt.Println(company3)
	fmt.Println()
	en := encryptions.GetPlus(60, 23)
	fmt.Println("Ans=", en)
	fmt.Println()

	eletricBill, waterBill, maintenanceBill := fileop.GetBill("6k", "4k", "12k")
	fmt.Println(eletricBill)
	fmt.Println(waterBill)
	fmt.Println(maintenanceBill)
	fmt.Println()

	file1, file2, file3 := files.GetFiles("officeFill", "directoryFile", "deviceFill")
	fmt.Println(file1)
	fmt.Println(file2)
	fmt.Println(file3)
	fmt.Println()

	items1, items2, items3, items4 := helper.GetCost("58k", "90k", "50k", "40k")
	fmt.Println(items1)
	fmt.Println(items2)
	fmt.Println(items3)
	fmt.Println(items4)

	goo := http.GetGo("GOOGLE")
	fmt.Println(goo)
	fmt.Println()
	type1, type2, type3, type4 := imageop.FetchImageType("JPEG", "PNG", "GIF", "PDF")
	fmt.Println(type1)
	fmt.Println(type2)
	fmt.Println(type3)
	fmt.Println(type4)
	fmt.Println()
	grade1, grade2, grade3 := jp.GetJps(82, 56, 94)
	fmt.Println(grade1)
	fmt.Println(grade2)
	fmt.Println(grade3)
	news := json.Fetchjson("Saturday break fast")
	fmt.Println(news)
	name1 := jsonop.GetPara("Liberty234")
	fmt.Println(name1)
	fmt.Println()
	logIn, logOut := logs.GetLogsTime("10am", "2pm")
	fmt.Println(logIn)
	fmt.Println(logOut)
	fmt.Println()
	net1, net2, net3 := network.GetNetworkType("MTN", "GLO", "Airtel")
	fmt.Println(net1)
	fmt.Println(net2)
	fmt.Println(net3)
	s := operation.GetCal(58, 34)
	fmt.Println("Ans:", s)
	fmt.Println()
	goods1, goods2, goods3 := parsers.GetGoodsCost("150k", "50k", "340k")
	fmt.Println(goods1)
	fmt.Println(goods2)
	fmt.Println(goods3)
	di := parsing.GetDivition(62, 12)
	fmt.Println("Ans=", di)
	fmt.Println()

	house1, house2, house3 := regexop.GetHouseType("Ranch", "Bungalow", "Cottage")
	fmt.Println(house1)
	fmt.Println(house2)
	fmt.Println(house3)
	web := saver.GetWeb("www.libohop.com")
	fmt.Println(web)
	t := saverops.GetTotal(65, 902, 32)
	fmt.Println("Ans:", t)

}
