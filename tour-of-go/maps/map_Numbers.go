package main

import "fmt"

func GetCountryNumber() {

	number := map[string]int{

		"india":    123,
		"USA":      874,
		"France":   864,
		"China":    903,
		"Pakistan": 534,
	}

	number["Russia"] = 645
	delete(number, "Pakistan")
	fmt.Println(number)
	fmt.Println(number["USA"])
	temp, ok := number["US"]
	fmt.Println(temp, ok)
	temp, ok = number["USA"]
	fmt.Println(temp)
	fmt.Println(len(number))

}
