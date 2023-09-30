package csv

import "fmt"

func GetName(FirstName, LastName, middleName string) (string, string, string) {
	FirstName = fmt.Sprintf("First Name: %s", FirstName)
	LastName = fmt.Sprintf("Last Name: %s", LastName)
	middleName = fmt.Sprintf("Middle Name: %s", middleName)
	return FirstName, LastName, middleName

}
