package sub

import "fmt"

func GetSubFun(sub string) (sub1 string) {
	sub1 = fmt.Sprintf("GB:%s", sub)
	return
}
