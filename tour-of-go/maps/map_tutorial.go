package main

import "fmt"

func GetItems() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	fmt.Println(mp)
	fmt.Println(mp["apple"])
	mp["grape"] = 8
	mp["apple"] = 4
	fmt.Println(mp)
	delete(mp, "grape")
	fmt.Println(mp)
	val, ok := mp["apple"]
	fmt.Println(val, ok)
	fmt.Println(len(mp))
}
