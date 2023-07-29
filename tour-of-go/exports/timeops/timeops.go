package timeops

import "fmt"

var Level = 100

func CarNumber() {
	carColor := []string{"blue", "red", "green"}
	carName := []string{"car"}
	for _, f := range carColor {
		for _, l := range carName {
			fmt.Println(f + " " + l)

		}

	}

}

func UserName() {
	fmt.Println("Signed By:", signed)
	user()
}
