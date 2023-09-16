package dp

import "fmt"

func GetAllName(FirstName1, LastName1, middleName1 string) (string, string, string) {
	fName := fmt.Sprintf("FirstName: %v", FirstName1)
	lName := fmt.Sprintf("LastName: %v", LastName1)
	mName := fmt.Sprintf("MiddleName: %v", middleName1)

	return fName, lName, mName

}
