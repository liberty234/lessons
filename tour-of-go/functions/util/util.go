package util

import "fmt"

func GetCarCost(carCost1, carCost int) (string, string) {
	amount1 := fmt.Sprintf("Benz:%v", carCost1)
	amount2 := fmt.Sprintf("Lexus:%v", carCost)
	return amount1, amount2

}
