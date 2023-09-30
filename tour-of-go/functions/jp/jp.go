package jp

import (
	"fmt"
)

func GetJps(grade1, grade2, grade3 int) (string, string, string) {
	score1 := fmt.Sprintf("First Grade:%v", grade1)
	score2 := fmt.Sprintf("Second Grade:%v", grade2)
	score3 := fmt.Sprintf("Third Grade:%v", grade3)
	return score1, score2, score3

}
