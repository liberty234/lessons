package main

import "fmt"

func GetInt() {

	for i := 1; i <= 1000; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
