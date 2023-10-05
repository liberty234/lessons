package main

import "fmt"

func FetchMap() {

	a1 := [5]int{1, 2, 3, 4, 5}
	a2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a1 == a2)

	a1 = [5]int{1, 2, 3, 4, 5}
	a2 = [5]int{1, 23, 43, 4, 5}
	fmt.Println(a1 == a2)

	m := map[string]int{}
	fmt.Println(m == nil)
	m["abc"] = 10
	val := m["abc"]
	fmt.Println(val)

	s := make(map[string]int, 10)
	fmt.Println(len(s))

}
