package auth

import "fmt"

func GetFullName(FirstName, LastName string) (string, string) {
	firstName := fmt.Sprintf("FirstName: %v\n", FirstName)
	lastName := fmt.Sprintf("LastName: %v", LastName)

	return firstName, lastName

}
