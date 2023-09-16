package json

import "fmt"

func Fetchjson(news string) string {
	js := fmt.Sprintf("TODAY NEWS:%s", news)
	return js

}
