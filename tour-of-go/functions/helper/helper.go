package helper

import "fmt"

func GetCost(items1, items2, items3, items4 string) (string, string, string, string) {
	items1 = fmt.Sprintf("Items1: %v", items1)
	items2 = fmt.Sprintf("Items2: %v", items2)
	items3 = fmt.Sprintf("Items3: %v", items3)
	items4 = fmt.Sprintf("Items4: %v", items4)

	return items1, items2, items3, items4

}
