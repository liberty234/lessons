package buffers

import "fmt"

var Level = 100

func Register() {
	firstName := []string{"Liberty", "Best", "Peace"}
	lastName := []string{"Ebikade"}
	for _, f := range firstName {
		for _, l := range lastName {
			fmt.Println(f + " " + l)

		}

	}

}

func Sub() {
	fmt.Println("SubB", subb)
	sub()
}
