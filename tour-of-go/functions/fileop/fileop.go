package fileop

import "fmt"

func GetBill(eletricBill, waterBill, maintenanceBill string) (eletricBill1 string, waterBill1 string, maintenanceBill1 string) {
	eletricBill1 = fmt.Sprintf("Electric Bill: %s", eletricBill)
	waterBill1 = fmt.Sprintf("Water Bill: %s", waterBill)
	maintenanceBill1 = fmt.Sprintf("Maintenance Bill: %s", maintenanceBill)
	return

}
