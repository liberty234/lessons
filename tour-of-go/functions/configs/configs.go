package configs

import "fmt"

func FetchNameVet(game1, game2, game3, game4 string) (string, string, string, string) {
	p := fmt.Sprintf("Game Name: %s", game1)
	g := fmt.Sprintf("Game Name: %s", game2)
	d := fmt.Sprintf("Game Name: %s", game3)
	a := fmt.Sprintf("Game Name: %s", game4)

	return p, g, d, a

}
