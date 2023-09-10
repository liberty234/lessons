package tags

import "fmt"

var Tags string = "This is my tags variable"

func GetTagFun() {
	a := "This is my function tags"
	fmt.Println("Tag:", a)

}
func GetTagsTypeFunc() {
	fmt.Println("tagsTypeVar:", tagsType)
	getTagsTypeFunc()

}
