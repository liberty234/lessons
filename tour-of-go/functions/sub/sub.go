package sub

import "fmt"

func GetSubFun(sub string) string {
	sub1 := fmt.Sprintf("GB:%s", sub)
	return sub1
}
