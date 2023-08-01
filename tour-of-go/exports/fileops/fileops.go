package fileops

import "fmt"

var PayBill string = "electric bill : 5k"

func GetBillPaid() {
	sd := 1000 + 5000
	fmt.Println("Total Bill Paid:", sd)

}

func GetBill() {
	totalBill := getBill(100, 50, 60)
	fmt.Println("Total Bill", totalBill)

	fmt.Println("Bill Amount:", amount)
}
