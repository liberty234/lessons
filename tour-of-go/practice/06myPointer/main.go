package main

import "fmt"

func main() {

	fmt.Println("welcome to class of pointer")
	var pt *int
	fmt.Println(pt)

	myNumber := 24

	var ptr = &myNumber

	fmt.Println("value of the pointer is ", ptr)
	fmt.Println("value of the pointer is ", *ptr)

	*ptr = 6 + *ptr
	fmt.Println("New value ", myNumber)
	fmt.Println()

	use1 := user2{
		FirstName: "Liberty",
		SurName:   "Ebikade",
		Age:       22,
	}
	fmt.Println(use1)

	user2, err := newUser(use1)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Get First Name:", user2.FirstName)
		user2.UpdateName("James")
		fmt.Println(user2)
	}

	Pointer()

	Pi := P{
		name: "lee",
		pin:  996700,
	}

	fmt.Println(&Pi)

}

type user1 struct {
	FirstName string
	SurName   string
	Age       int
}

type user2 struct {
	FirstName string
	SurName   string
	Age       int
}

func newUser(n user2) (*user1, error) {
	user1 := user1{FirstName: n.FirstName}
	user1.SurName = n.SurName
	user1.Age = n.Age

	return &user1, nil

}

func (u user1) GetName() string {
	return u.FirstName

}

func (u *user1) UpdateName(FirstName string) *user1 {
	u.FirstName = FirstName
	return u

}

type P struct {
	name string
	pin  int
}

func Pointer() {

	fmt.Println("welcome to pointer leasons")
	myAge := 22
	var prt = &myAge

	fmt.Println("Value address", &prt)
	fmt.Println("Value address", *prt)

	*prt = *prt + 1
	fmt.Println("New Age:", myAge)

}
