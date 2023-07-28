package score

import "fmt"

var MainScore = "90%"

func Source() {
	bestscore := 85
	main := 90

	fmt.Println("SourceScore", main-bestscore)

}

func TotalScore() {
	fmt.Println("Best score:", bestscore)
	totalSocre()
}
