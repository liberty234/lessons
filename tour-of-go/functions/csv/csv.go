package csv

import "fmt"

func GetName(FirstName, LastName, middleName string) (string, string, string) {
	f := fmt.Sprintf("First Name: %s", FirstName)
	l := fmt.Sprintf("Last Name: %s", LastName)
	m := fmt.Sprintf("Middle Name: %s", middleName)
	return f, l, m

}
