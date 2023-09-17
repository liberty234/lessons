package webs

import "fmt"

func GetPost(gooo, roo string) (string, string) {

	d := fmt.Sprintf("Post%s", gooo)
	s := fmt.Sprintf("Post%s", roo)
	return d, s

}
