package deck

import "fmt"

func newDeck(red, blue int) int {
	return red + blue
}

func Cards() {
	Cards := []string{"five of diamond", "Ace of Spade", "five of Spade"}
	Cards = append(Cards, "Ace of diamond")
	for i, card := range Cards {
		fmt.Println(i, card)
	}
}
