package xmlop

import "fmt"

func CreateAddress(street, city, state string, postalCode int) (string, string, string, string) {
	street1 := fmt.Sprintf("Street:%s", street)
	city1 := fmt.Sprintf("City:%s", city)
	state1 := fmt.Sprintf("State:%s", state)
	postalCode1 := fmt.Sprintf("PostCode:%v", postalCode)
	return street1, city1, state1, postalCode1
}
