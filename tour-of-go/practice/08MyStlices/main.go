package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slice in go toturial")

	var fruitList = []string{
		"mango",
		"tomato",
		"Orange",
		"pineapple",
		"Apple",
	}

	fmt.Println("FruitList:", fruitList)

	fruitList = append(fruitList, "grape", "cashew")
	fmt.Println("FruitList :", fruitList)
	fmt.Println("FruitList:", len(fruitList))

	fruitList = append(fruitList[1:4])
	fmt.Println("Full fruit List :", fruitList)
	fmt.Println("FruitList:", len(fruitList))

	myScore := make([]int, 4)

	myScore[0] = 485
	myScore[1] = 748
	myScore[2] = 894
	myScore[3] = 938
	//myScore[4] = 785

	myScore = append(myScore, 223, 900, 435)

	fmt.Println("MyScore:", myScore)
	sort.Ints(myScore)
	fmt.Println("mySocre:", myScore)

	courses := []string{"python", "Go", "Ruby", "javascript", "reactjc", "c++"}
	fmt.Println(courses)

}
