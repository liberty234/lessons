package fileops

import (
	"fmt"
)

func Bill() {
	totalBill := bill(100, 50, 60)
	fmt.Println("Total Bill", totalBill)
}
