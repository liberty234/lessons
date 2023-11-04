package main

import "fmt"

func GetMap() {

	//arrey

	/*var arr [5]int

	arr[1] = 30
	arr[4] = 12
	arr[0] = 90

	fmt.Println(arr[0])*/

	//slice
	s := make([]int, 5)

	s = append(s, 80, 69, 54, 23, 42)

	cpl := make([]int, len(s))

	cpl[4] = 45
	fmt.Println(cpl)
	s[3] = 34

	fmt.Println(s)

	//map

	m1 := make(map[string]int)
	m1["key1"] = 35
	m1["key2"] = 64
	m1["some other keys"] = 29

	fmt.Println(m1["key1"])

	fmt.Println(m1)

	value, present := m1["NON-Existent"]
	fmt.Println(value, present)

}
