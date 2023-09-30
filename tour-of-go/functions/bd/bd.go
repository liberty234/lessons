package bd

import "fmt"

func GetPersonName(name string) string {
	personName := fmt.Sprintf("Full Name:%s", name)
	return personName
}
