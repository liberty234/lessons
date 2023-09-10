package email

import "fmt"

var Email string = "libertyebikade06@gmail.com"

func GetEmailFunc() {
	e := 51
	fmt.Printf("Notification: you have %v email notifications", e)

}

func GetValidEmail() {
	fmt.Println("\nEnter comfirmation email:", email)
	getEmailFunc()

}
