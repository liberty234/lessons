package bac

import "fmt"

func GetClass(class1, class2, class3, class4 string) (string, string, string, string) {
	a := fmt.Sprintf("Class1 %v", class1)
	b := fmt.Sprintf("Class2 %v", class2)
	c := fmt.Sprintf("Class2 %v", class3)
	d := fmt.Sprintf("Class2 %v", class4)

	return a, b, c, d

}
