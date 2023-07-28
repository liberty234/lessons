package logs

import "fmt"

var LogTime = "i log in 9:30 today"

func LogName() {
	log := []string{"Liberty", "Best"}
	fmt.Println("LogName:", log)

}

func UserName() {
	fmt.Println("Signed By:", signed)
	user()
}
