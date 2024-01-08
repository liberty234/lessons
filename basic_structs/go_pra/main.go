package main

import "fmt"

func main() {

	user1 := newUser1{
		firstName:  "Liberty",
		lastName:   "Ebiakde",
		complexion: complexion2,
		age:        22,
		email:      "liberty06@gmail.com",
		hasGlasses: false,
		hasCar:     false,
	}

	user2, err := newUser(user1)

	if err != nil {
		fmt.Println("Errors", err.Error())
	}

	if user2 != nil {
		fmt.Println(user1)
		fmt.Println("firstName:", user1.firstName)
		fmt.Println("lastName:", user1.lastName)
		fmt.Println("")
		user2.UpdateFirstName("Best")
		user2.UpdateLastName("Ebikade")
		user2.UpdateAge(18)
		user2.UpdateEmail("best234@gmail.com")
		fmt.Println(user2)
		fmt.Println("FirstName:", user2.firstName)
		fmt.Println("LastName:", user2.lastName)

	}

}
