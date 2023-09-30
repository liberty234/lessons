package bac

import "fmt"

func GetClass(class1, class2, class3, class4 string) (string, string, string, string) {
	class1 = fmt.Sprintf("Class1 %v", class1)
	class2 = fmt.Sprintf("Class2 %v", class2)
	class3 = fmt.Sprintf("Class2 %v", class3)
	class4 = fmt.Sprintf("Class2 %v", class4)

	return class1, class2, class3, class4

}
