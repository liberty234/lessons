package http

import "fmt"

var log string = "log in"

func getWebsite() {
	w := []string{"Facebook", "WhatsApp", "twitter"}
	for _, v := range w {
		fmt.Println(v)

	}

}
