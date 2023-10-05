package books

import "fmt"

func FetchBooked(booked1, booked2, booked3, booked4, booked5, booked6 int) (string, string, string, string, string, string) {
	booked := fmt.Sprintf("Ticket booked:%v", booked1)
	booked12 := fmt.Sprintf("Ticket booked:%v", booked2)
	booked13 := fmt.Sprintf("Ticket booked:%v", booked3)
	booked14 := fmt.Sprintf("Ticket booked:%v", booked4)
	booked15 := fmt.Sprintf("Ticket booked:%v", booked5)
	booked16 := fmt.Sprintf("Ticket booked:%v", booked6)
	return booked, booked12, booked13, booked14, booked15, booked16
}
