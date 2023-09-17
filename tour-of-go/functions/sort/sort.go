package sort

import "fmt"

func GetVex(vex1, vex2, vex3 string) (string, string, string) {
	v1 := fmt.Sprintf("Event1:%s", vex1)
	v2 := fmt.Sprintf("Event2:%s", vex2)
	v3 := fmt.Sprintf("Event3:%s", vex3)
	return v1, v2, v3
}
