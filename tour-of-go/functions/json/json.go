package json

import "fmt"

func Fetchjson(news string) string {
	breakNews := fmt.Sprintf("TODAY NEWS:%s", news)
	return breakNews

}
