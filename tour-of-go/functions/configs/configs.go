package configs

import "fmt"

func FetchNameVet(game1, game2, game3, game4 string) (gameplay1 string, gameplay2 string, gameplay3 string, gameplay4 string) {
	gameplay1 = fmt.Sprintf("Game Name: %s", game1)
	gameplay2 = fmt.Sprintf("Game Name: %s", game2)
	gameplay3 = fmt.Sprintf("Game Name: %s", game3)
	gameplay4 = fmt.Sprintf("Game Name: %s", game4)

	return

}
