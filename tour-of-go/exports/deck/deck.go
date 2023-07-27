package deck

import "fmt"

var seed string = "beans seed"

var MyName string = "my name is Liberty"

func Cards() {
	Cards := []string{"five of diamond", "Ace of Spade", "five of Spade"}
	Cards = append(Cards, "Ace of diamond")
	for i, card := range Cards {
		fmt.Println(i, card)
	}

}

func newDeck() {
	dek := 54 * 34
	fmt.Println("dek", dek)
}
