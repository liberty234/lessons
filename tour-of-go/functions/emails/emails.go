package emails

import "fmt"

func FetchCompanyEmail(company1, company2, company3 string) (string, string, string) {
	c1 := fmt.Sprintf("Best Company Email Address: %s", company1)
	c2 := fmt.Sprintf("Godbless Company Email Address: %s", company2)
	c3 := fmt.Sprintf("Liberty Company Email Address: %s", company3)

	return c1, c2, c3

}
