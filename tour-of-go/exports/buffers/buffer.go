package buffers

import "fmt"

var Level int = 100

func GetRegister() {
	firstName := []string{"Liberty", "Best", "Peace"}
	lastName := []string{"Ebikade"}
	for _, f := range firstName {
		for _, l := range lastName {
			fmt.Println(f + " " + l)

		}

	}

}

func GetSub() {
	fmt.Println("SubB", subb)
	getSub()
}
