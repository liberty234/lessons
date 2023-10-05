package bac

import "fmt"

func GetClass(class1, class2, class3, class4 string) (class5 string, class6 string, class7 string, class8 string) {
	class5 = fmt.Sprintf("Class1 %v", class1)
	class6 = fmt.Sprintf("Class2 %v", class2)
	class7 = fmt.Sprintf("Class2 %v", class3)
	class8 = fmt.Sprintf("Class2 %v", class4)

	return

}
