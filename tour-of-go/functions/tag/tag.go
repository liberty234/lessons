package tag

import "fmt"

func GetTagFun(tag string) string {
	tag = fmt.Sprintf("Tags:%s", tag)
	return tag
}
