package deck

import "fmt"

var MyName string = "my name is Liberty"

func GetCards() {
	Cards := []string{"five of diamond", "Ace of Spade", "five of Spade"}
	Cards = append(Cards, "Ace of diamond")
	for i, card := range Cards {
		fmt.Println(i, card)
	}

}
func FetchNewDeck() {
	fmt.Println(seed)
	getNewDeck()

}
