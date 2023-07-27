package fileops

import "fmt"

var Paybill string = "electric bill : 5k"

func Bill() {
	totalBill := bill(100, 50, 60)
	fmt.Println("Total Bill", totalBill)
}
