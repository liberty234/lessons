package sorting

import "fmt"

func GetSort(eye1, eye2, eye3 string) (string, string, string) {
	eye1 = fmt.Sprintf("Eye Color:%s", eye1)
	eye2 = fmt.Sprintf("Eye Color:%s", eye2)
	eye3 = fmt.Sprintf("Eye Color:%s", eye3)
	return eye1, eye2, eye3
}
