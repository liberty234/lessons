package http

import "fmt"

func GetGo(goo string) string {
	v := fmt.Sprintf("Browser: %s", goo)
	return v

}
