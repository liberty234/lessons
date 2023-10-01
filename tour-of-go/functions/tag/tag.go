package tag

import "fmt"

func GetTagFun(tag string) (tag1 string) {
	tag1 = fmt.Sprintf("Tags:%s", tag)
	return
}
