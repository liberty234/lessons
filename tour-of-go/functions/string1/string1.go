package string1

import "fmt"

func GetPersonalInfo(FirstName, LastName, middleName string) (string, string, string) {
	firstName := fmt.Sprintf("First Name: %s", FirstName)
	lastName := fmt.Sprintf("Last Name: %s", LastName)
	MiddleName := fmt.Sprintf("Middle Name: %s", middleName)
	return firstName, lastName, MiddleName

}
