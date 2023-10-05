package jp

import (
	"fmt"
)

func GetJps(grade1, grade2, grade3 int) (score11 string, score22 string, score33 string) {
	score11 = fmt.Sprintf("First Grade:%v", grade1)
	score22 = fmt.Sprintf("Second Grade:%v", grade2)
	score33 = fmt.Sprintf("Third Grade:%v", grade3)
	return

}
