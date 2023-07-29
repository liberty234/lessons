package booking

import "fmt"

var Booking = "Four people has been booked today"

func Booked() {
	bookedName := []string{"Liberty", "Favor", "Gift", "Grace"}
	for _, v := range bookedName {
		fmt.Println(v)

	}
}
func TicketBooked() {
	fmt.Println("VIP:", booking)
	bookedTicket()

}
