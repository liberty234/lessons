package webs

import "fmt"

func GetPost(google, room string) (string, string) {

	like1 := fmt.Sprintf("Post%s", google)
	like2 := fmt.Sprintf("Post%s", room)
	return like1, like2

}
