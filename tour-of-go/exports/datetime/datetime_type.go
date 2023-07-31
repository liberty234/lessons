package datetime

import "fmt"

var week string = "Today is the 6th day of the week"

func getWeeks() {
	s := "sunday"
	m := "monday"
	t := "tuesday"
	w := "wednesday"
	th := "thursday"
	f := "friday"
	sa := "saturday"

	fmt.Println("Day Of The Weeks:", s, m, t, w, th, f, sa)
}
