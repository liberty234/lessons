package deck

import "fmt"

func GetCard(ace, five, four, two string) (string, string, string, string) {
	deck1 := fmt.Sprintf("Ace %s", ace)
	deck2 := fmt.Sprintf("Five %s", five)
	deck3 := fmt.Sprintf("Four %s", four)
	deck4 := fmt.Sprintf("Two %s", two)
	return deck1, deck2, deck3, deck4

}
