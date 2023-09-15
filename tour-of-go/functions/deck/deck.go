package deck

import "fmt"

func GetCard(ace, five, four, two string) (string, string, string, string) {
	A := fmt.Sprintf("Ace %s", ace)
	F := fmt.Sprintf("Five %s", five)
	Fo := fmt.Sprintf("Four %s", four)
	T := fmt.Sprintf("Two %s", two)
	return A, F, Fo, T

}
