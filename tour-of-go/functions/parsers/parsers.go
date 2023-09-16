package parsers

import "fmt"

func GetGoodsCost(goods1, goods2, goods3 string) (string, string, string) {
	b := fmt.Sprintf("Bread:%s", goods1)
	be := fmt.Sprintf("Beans:%s", goods2)
	r := fmt.Sprintf("Rice:%s", goods3)
	return b, be, r

}
