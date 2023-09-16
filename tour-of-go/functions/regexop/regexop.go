package regexop

import "fmt"

func GetHouseType(house1, house2, house3 string) (string, string, string) {
	m := fmt.Sprintf("House1:%s", house1)
	g := fmt.Sprintf("House2:%s", house2)
	a := fmt.Sprintf("House3:%s", house3)
	return m, g, a
}
