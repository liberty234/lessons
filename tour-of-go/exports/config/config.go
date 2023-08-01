package config

import "fmt"

var GameName string = "FiFa 2023"

func FetchNameVet() {
	Namevex := []string{"boxing game", "dice game", "Ludo game"}
	Namevex = append(Namevex, "Spider man game")
	for i, car := range Namevex {
		fmt.Println(i, car)
	}

}

func FetchCalcu() {
	fmt.Println("Sub total:", sub)
	getCal()

}
