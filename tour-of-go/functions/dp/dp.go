package dp

import "fmt"

func GetAllName(FirstName1, LastName1, middleName1 string) (string, string, string) {
	fullName := fmt.Sprintf("FirstName: %v", FirstName1)
	lastName := fmt.Sprintf("LastName: %v", LastName1)
	middleName := fmt.Sprintf("MiddleName: %v", middleName1)

	return fullName, lastName, middleName

}
