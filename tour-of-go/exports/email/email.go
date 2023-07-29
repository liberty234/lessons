package email

import "fmt"

var Email = "libertyebikade06@gmail.com"

func Emailfunc() {
	e := 51
	fmt.Printf("Notification: you have %v email notifications", e)

}

func ValidEmail() {
	fmt.Println("\nEnter comfirmation email:", email)
	emailfunc()

}
