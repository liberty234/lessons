package statistics

import "fmt"

var stat float32 = 25.5

func getStatic() {
	firstScore := 202
	secondScore := 262
	divide := 4

	fmt.Println("TotalScore:", firstScore+secondScore/divide)

}
