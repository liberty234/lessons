/*package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}*/

package main

import "fmt"

func add(x, y int) int {
	return y + x

}

func main() {

	result := add(32, 8)

	fmt.Println(result)
}
