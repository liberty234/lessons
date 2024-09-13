package main

import "errors"

const (
	complexion1 = "Black"
	complexion2 = "Brown"
	complexion3 = "white"
)

var (
	complexions = map[string]string{
		complexion1: complexion1,
		complexion2: complexion2,
		complexion3: complexion3,
	}
)

type user struct {
	firstName  string
	lastName   string
	complexion string
	age        int
	email      string
	hasGlasses bool
	hasCar     bool
}

type newUser1 struct {
	firstName  string
	lastName   string
	complexion string
	age        int
	email      string
	hasGlasses bool
	hasCar     bool
}

func newUser(user1 newUser1) (*user, error) {
	user := user{firstName: user1.firstName}
	user.lastName = user1.lastName
	user.complexion = user1.complexion
	user.age = user1.age
	user.email = user1.email
	user.hasGlasses = user1.hasGlasses
	user.hasCar = user1.hasCar

	if _, complexionIsFound := complexions[user.complexion]; !complexionIsFound {
		return nil, errors.New("complexion not found")
	}
	return &user, nil

}

func (u *user) GetFirstName() string {
	return u.firstName
}

func (u *user) GetLastName() string {
	return u.lastName

}

func (u *user) GetComplexion() string {
	return u.complexion

}

func (u *user) GetAge() int {
	return u.age

}

func (u *user) GetEmail() string {
	return u.email

}

func (u *user) UpdateFirstName(firstName string) *user {
	u.firstName = firstName
	return u

}

func (u *user) UpdateLastName(lastName string) *user {
	u.lastName = lastName
	return u

}
func (u *user) UpdateComplexion(complexion string) *user {
	u.complexion = complexion
	return u

}
func (u *user) UpdateAge(age int) *user {
	u.age = age
	return u

}
func (u *user) UpdateEmail(email string) *user {
	u.email = email
	return u

}
