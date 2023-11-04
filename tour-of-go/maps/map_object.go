package main

import "fmt"

func FetchObject() {

	var m map[string]string
	fmt.Println(m == nil)
	m = map[string]string{}
	fmt.Println(m == nil)
	fmt.Println(len(m))
	m = make(map[string]string, 65)
	fmt.Println(len(m))
	m = map[string]string{"Liberty": "Programmer"}
	fmt.Println(m)
	fmt.Println(len(m))
	m["hanny"] = "Not a programmer"
	fmt.Println(m)
	fmt.Println(len(m))
	m["hanny"] = "a programmer to be"
	fmt.Println(m)
	delete(m, "hanny")
	m["Liberty"] = "baller"
	fmt.Println(m)

	for name, title := range m {
		fmt.Println(name, title)

	}

	title, ok := m["hanny"]

	if ok {
		fmt.Println(title)
	} else {
		fmt.Println("Not found")
	}

}
