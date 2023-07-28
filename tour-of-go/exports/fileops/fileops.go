package fileops

import "fmt"

var Paybill string = "electric bill : 5k"

func Billpaid() {
	sd := 1000 + 5000
	fmt.Println("Total Bill Paid:", sd)

}

func Bill() {
	totalBill := bill(100, 50, 60)
	fmt.Println("Total Bill", totalBill)

	fmt.Println("Bill Amount:", amount)
}
