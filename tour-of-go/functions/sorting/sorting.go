package sorting

import "fmt"

func GetSort(eye1, eye2, eye3 string) (string, string, string) {
	e1 := fmt.Sprintf("Eye Color:%s", eye1)
	e2 := fmt.Sprintf("Eye Color:%s", eye2)
	e3 := fmt.Sprintf("Eye Color:%s", eye3)
	return e1, e2, e3
}
