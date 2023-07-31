package state

import "fmt"

var StateOfOrigin string = "EDO STATE"

func GetStateName() {
	e := "EDO STATE"
	a := "ABIA STATE"
	l := "LAGOS STATE"
	an := "ANAMBRA STATE"
	fmt.Println("Some state in Nigeria:", e, a, l, an)

}
func GetOrigin() {
	fmt.Println("Place:", state)
	getOrigin()

}
