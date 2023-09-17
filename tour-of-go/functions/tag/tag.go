package tag

import "fmt"

func GetTagFun(tag string) string {
	c := fmt.Sprintf("Tags:%s", tag)
	return c
}
