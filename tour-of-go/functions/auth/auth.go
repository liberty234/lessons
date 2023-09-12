package auth

import "fmt"

func GetFullName(FirstName, LastName string) (string, string) {
	fName := fmt.Sprintf("FirstName: %v\n", FirstName)
	lName := fmt.Sprintf("LastName: %v", LastName)

	return fName, lName

}
