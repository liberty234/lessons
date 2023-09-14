package bufa

import "fmt"

func GetRegister(reg1, reg2, reg3, reg4 string) (string, string, string, string) {
	el := fmt.Sprintf("Reg1:%v", reg1)
	ep := fmt.Sprintf("Reg1:%v", reg2)
	et := fmt.Sprintf("Reg1:%v", reg3)
	eu := fmt.Sprintf("Reg1:%v", reg4)
	return el, ep, et, eu

}
