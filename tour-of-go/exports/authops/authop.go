package authops

import "fmt"

func maths(x int, y int, z int, k int) int {
	return x + y/z + k

}

func Premiar() {
	NameLeague := []string{"Premiar", "laliga", "seria "}
	postLeague := []string{"League"}
	for _, n := range NameLeague {
		for _, p := range postLeague {
			fmt.Println(n + " " + p)

		}

	}

}
