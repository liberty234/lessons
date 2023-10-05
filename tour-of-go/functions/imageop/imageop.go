package imageop

import "fmt"

func FetchImageType(type1, type2, type3, type4 string) (string, string, string, string) {
	file1 := fmt.Sprintf("Type1:%s", type1)
	file2 := fmt.Sprintf("Type2:%s", type2)
	file3 := fmt.Sprintf("Type3:%s", type3)
	file4 := fmt.Sprintf("Type4:%s", type4)
	return file1, file2, file3, file4

}
