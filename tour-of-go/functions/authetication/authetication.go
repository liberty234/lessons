package authetication

import "fmt"

func GetPersonName(Name string) string {
	pName := fmt.Sprintf("First Person %s", Name)
	return pName
}
