package score

import "fmt"

var MainScore string = "90%"

func GetSource() {
	bestscore := 85
	main := 90

	fmt.Println("SourceScore", main-bestscore)

}

func GetTotalScore() {
	fmt.Println("Best score:", bestscore)
	getTotalSocre()
}
