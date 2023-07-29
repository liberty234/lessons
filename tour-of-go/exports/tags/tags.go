package tags

import "fmt"

var Tags = "This is my tags variable"

func Tagfun() {
	a := "This is my function tags"
	fmt.Println("Tag:", a)

}
func TagsTypeFunc() {
	fmt.Println("tagsTypeVar:", tagsType)
	tagsTypefunc()

}
