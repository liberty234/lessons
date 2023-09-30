package authetication

import "fmt"

func GetPersonName(Name string) string {
	personName := fmt.Sprintf("First Person %s", Name)
	return personName
}
