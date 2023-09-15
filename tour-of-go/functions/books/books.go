package books

import "fmt"

func FetchBooked(booked1, booked2, booked3, booked4, booked5, booked6 int) (string, string, string, string, string, string) {
	f := fmt.Sprintf("Ticket booked:%v", booked1)
	h := fmt.Sprintf("Ticket booked:%v", booked2)
	e := fmt.Sprintf("Ticket booked:%v", booked3)
	g := fmt.Sprintf("Ticket booked:%v", booked4)
	t := fmt.Sprintf("Ticket booked:%v", booked5)
	r := fmt.Sprintf("Ticket booked:%v", booked6)
	return f, h, e, g, t, r
}
