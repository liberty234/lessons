package bd

import "fmt"

func GetPersonName(name string) (personName string) {
	personName = fmt.Sprintf("Full Name:%s", name)
	return
}
