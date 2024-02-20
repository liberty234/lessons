package main

import "fmt"

func main() {
	fmt.Println("welcome to loop in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "friday", "saturday"}
	fmt.Println("days of the week:", days)
	//for _, value := range days {
	//	fmt.Println("days of the week:", value)
	//}

	for i := range days {
		fmt.Println(days[i])
	}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

}
