package jp

import "fmt"

func GetJps(grade1, grade2, grade3 int) (string, string, string) {
	g1 := fmt.Sprintf("First Grade:%v", grade1)
	g2 := fmt.Sprintf("Second Grade:%v", grade2)
	g3 := fmt.Sprintf("Third Grade:%v", grade3)
	return g1, g2, g3

}
