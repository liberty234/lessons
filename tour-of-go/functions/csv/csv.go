package csv

import "fmt"

func GetName(FirstName, LastName, middleName string) (FirstName1 string, LastName2 string, middleName3 string) {
	FirstName1 = fmt.Sprintf("First Name: %s", FirstName)
	LastName2 = fmt.Sprintf("Last Name: %s", LastName)
	middleName3 = fmt.Sprintf("Middle Name: %s", middleName)
	return FirstName, LastName, middleName

}
