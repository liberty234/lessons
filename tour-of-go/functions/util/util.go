package util

import "fmt"

func GetCarCost(carCost1, carCost int) (string, string) {
	f := fmt.Sprintf("Benz:%v", carCost1)
	e := fmt.Sprintf("Lexus:%v", carCost)
	return f, e

}
