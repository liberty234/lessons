package fileop

import "fmt"

func GetBill(eletricBill, waterBill, maintenanceBill string) (string, string, string) {
	eb := fmt.Sprintf("Electric Bill: %s", eletricBill)
	wb := fmt.Sprintf("Water Bill: %s", waterBill)
	mb := fmt.Sprintf("Maintenance Bill: %s", maintenanceBill)
	return eb, wb, mb

}
