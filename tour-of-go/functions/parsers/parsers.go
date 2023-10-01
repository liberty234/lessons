package parsers

import "fmt"

func GetGoodsCost(goods1, goods2, goods3 string) (firstGoods string, secondGoods string, thirdgoods string) {
	firstGoods = fmt.Sprintf("Bread:%s", goods1)
	secondGoods = fmt.Sprintf("Beans:%s", goods2)
	thirdgoods = fmt.Sprintf("Rice:%s", goods3)
	return

}
