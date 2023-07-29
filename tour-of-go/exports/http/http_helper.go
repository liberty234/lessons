package http

import "fmt"

var log = "log in"

func website() {
	w := []string{"Facebook", "WhatsApp", "tweeter"}
	for _, v := range w {
		fmt.Println(v)

	}

}
