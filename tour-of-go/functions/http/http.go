package http

import "fmt"

func GetGo(goo string) string {
	value := fmt.Sprintf("Browser: %s", goo)
	return value

}
