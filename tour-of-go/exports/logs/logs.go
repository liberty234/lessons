package logs

import "fmt"

var LogTime string = "i log in 9:30 today"

func GetLogName() {
	log := []string{"Liberty", "Best"}
	fmt.Println("LogName:", log)

}

func GetUserName() {
	fmt.Println("Signed By:", signed)
	getUser()
}
