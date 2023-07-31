package datetime

import (
	"fmt"
	"time"
)

var TimeDate string = "10:17-28-07-2023"

func GetDateTime() {
	today := "friday 28 july"
	fmt.Println("Today date:", time.Now(), today)
}

func GetWeeks() {
	fmt.Println("Weeks name:", week)
	getWeeks()

}
