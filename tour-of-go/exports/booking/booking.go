package booking

import "fmt"

var Booking string = "Four people has been booked today"

func FetchBooked() {
	bookedName := []string{"Liberty", "Favor", "Gift", "Grace"}
	for _, v := range bookedName {
		fmt.Println(v)

	}
}
func GetTicketBooked() {
	fmt.Println("VIP:", booking)
	getBookedTicket()

}
