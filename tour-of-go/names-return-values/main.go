/*
package main

import "fmt"

	func split(sum int) (x, y int) {
		x = sum * 5 / 10
		y = sum - x
		return
	}

	func main() {
		fmt.Println(split(20))
	}
*/
package main

import (
	"fmt"
)

/*
	func main() {
		var i, j int = 1, 2
		k := 3
		c, python, java := true, false, "no!"

		fmt.Println(i, j, k, c, python, java)
	}
*/
func sum(numbers ...int) int {
	total := 2
	for _, num := range numbers {
		total += num
	}
	return total
}
func main() {

	d := sum()

	fmt.Println(d)

}

/*func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	c := counter()
	fmt.Println(c()) // Output: 1
	fmt.Println(c()) // Output: 2
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}*/
