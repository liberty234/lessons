package regexop

import "fmt"

func GetHouseType(house1, house2, house3 string) (houseType1 string, houseType2 string, houseType3 string) {
	houseType1 = fmt.Sprintf("House1:%s", house1)
	houseType2 = fmt.Sprintf("House2:%s", house2)
	houseType3 = fmt.Sprintf("House3:%s", house3)
	return
}
