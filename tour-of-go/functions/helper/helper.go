package helper

import "fmt"

func GetCost(items1, items2, items3, items4 string) (string, string, string, string) {
	i1 := fmt.Sprintf("Items1: %v", items1)
	i2 := fmt.Sprintf("Items2: %v", items2)
	i3 := fmt.Sprintf("Items3: %v", items3)
	i4 := fmt.Sprintf("Items4: %v", items4)

	return i1, i2, i3, i4

}
