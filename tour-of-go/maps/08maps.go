package main

import "fmt"

func GetScore() {

	score := make(map[string]int)
	score["josh"] = 78
	score["ron"] = 84
	score["sam"] = 92
	score["vicky"] = 68

	fmt.Println(score)
	getRonScore := score["ron"]
	fmt.Println(getRonScore)
	delete(score, "vicky")

	fmt.Println(score)
	for k, v := range score {
		fmt.Printf("score of %v is %v\n", k, v)
	}

}
