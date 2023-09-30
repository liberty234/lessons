package fileop

import "fmt"

func GetBill(eletricBill, waterBill, maintenanceBill string) (string, string, string) {
	eletricBill = fmt.Sprintf("Electric Bill: %s", eletricBill)
	waterBill = fmt.Sprintf("Water Bill: %s", waterBill)
	maintenanceBill = fmt.Sprintf("Maintenance Bill: %s", maintenanceBill)
	return eletricBill, waterBill, maintenanceBill

}
