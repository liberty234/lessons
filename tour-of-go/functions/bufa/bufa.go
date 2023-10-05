package bufa

import "fmt"

func GetRegister(reg1, reg2, reg3, reg4 string) (string, string, string, string) {
	fullName1 := fmt.Sprintf("Reg1:%v", reg1)
	fullName2 := fmt.Sprintf("Reg1:%v", reg2)
	fullName3 := fmt.Sprintf("Reg1:%v", reg3)
	fullName4 := fmt.Sprintf("Reg1:%v", reg4)
	return fullName1, fullName2, fullName3, fullName4

}
