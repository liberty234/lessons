package timeops

import "fmt"

var Level int = 100

func GetCarNumber() {
	carColor := []string{"blue", "red", "green"}
	carName := []string{"car"}
	for _, f := range carColor {
		for _, l := range carName {
			fmt.Println(f + " " + l)

		}

	}

}

func GetUserName() {
	fmt.Println("Signed By:", signed)
	getUser()
}
