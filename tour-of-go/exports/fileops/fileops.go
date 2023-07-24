package fileops

func bill(foodstuff int, fees int, gas int) int {
	return foodstuff + fees + gas

}

type Paybill struct {
	Fees   string
	ElBill string
	Gas    string
}
