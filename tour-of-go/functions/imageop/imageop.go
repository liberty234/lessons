package imageop

import "fmt"

func FetchImageType(type1, type2, type3, type4 string) (string, string, string, string) {
	t1 := fmt.Sprintf("Type1:%s", type1)
	t2 := fmt.Sprintf("Type2:%s", type2)
	t3 := fmt.Sprintf("Type3:%s", type3)
	t4 := fmt.Sprintf("Type4:%s", type4)
	return t1, t2, t3, t4

}
