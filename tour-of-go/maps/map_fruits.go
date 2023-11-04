package main

import "fmt"

func GetFruits() {
	m := make(map[string][]string)

	m["fruits"] = []string{"Apple", "pear", "cherry", "orange"}
	m["vegetables"] = []string{"Cucumber", "pumpkin", "yam", "potato"}

	fmt.Println(m)
	fruits := m["fruits"]
	fmt.Println("get fruits", fruits)

	delete(m, "fruits")
	fmt.Println(m)

	for k, v := range m {

		fmt.Println(k, v)

	}
	fmt.Println(len(m))

}
