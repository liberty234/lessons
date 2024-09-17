package main

import "fmt"

func main() {
	classNameNumber := map[string]int{
		"Ebikade": 4,
		"Oba":     6,
		"Grace":   8,
		"Hope":    2,
	}

	classNameNumber["Gift"] = 7
	for k, v := range classNameNumber {
		fmt.Println(k, v)

	}
	fmt.Println(len(classNameNumber))
	GetCountryPopulation()
	GetScore()
	GetItems()
	GetAges()
	GetMenu()
	FetchLanguage()
	FetchAges()
	GetPizza()
	GetFruits()
	FetchObject()
	GetCountryNumber()
	GetMap()
	GetMonthSales()
	GetNation()
	FetchMap()

}
