package xmlop

import "fmt"

func CreateAddress(street, city, state string, postalCode int) (string, string, string, string) {
	x := fmt.Sprintf("Street:%s", street)
	c := fmt.Sprintf("City:%s", city)
	s := fmt.Sprintf("State:%s", state)
	p := fmt.Sprintf("PostCode:%v", postalCode)
	return x, c, s, p
}
