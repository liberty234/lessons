package buffers

import "fmt"

func Register() {
	firstName := []string{"Liberty", "Best", "Peace"}
	lastName := []string{"Ebikade"}
	for _, f := range firstName {
		for _, l := range lastName {
			fmt.Println(f + " " + l)

		}

	}

}
func tans(x, y int) int {
	return x * y
}
