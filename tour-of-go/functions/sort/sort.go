package sort

import "fmt"

func GetVex(vex1, vex2, vex3 string) (string, string, string) {
	vex1 = fmt.Sprintf("Event1:%s", vex1)
	vex2 = fmt.Sprintf("Event2:%s", vex2)
	vex3 = fmt.Sprintf("Event3:%s", vex3)
	return vex1, vex2, vex3
}
