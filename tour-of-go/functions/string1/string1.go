package string1

import "fmt"

func GetPersonalInfo(FirstName, LastName, middleName string) (firstName1 string, lastName1 string, MiddleName1 string) {
	firstName1 = fmt.Sprintf("First Name: %s", FirstName)
	lastName1 = fmt.Sprintf("Last Name: %s", LastName)
	MiddleName1 = fmt.Sprintf("Middle Name: %s", middleName)
	return
}
