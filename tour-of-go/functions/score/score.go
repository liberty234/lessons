package score

import "fmt"

func GetSources(score1, score2 int) (total1 string, total2 string) {
	total1 = fmt.Sprintf("Score:%v", score1)
	total2 = fmt.Sprintf("Score:%v", score2)
	return

}
