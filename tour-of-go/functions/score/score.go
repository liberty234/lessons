package score

import "fmt"

func GetSources(score1, score2 int) (string, string) {
	s1 := fmt.Sprintf("Score:%v", score1)
	s2 := fmt.Sprintf("Score:%v", score2)
	return s1, s2

}
