package emails

import "fmt"

func FetchCompanyEmail(company1, company2, company3 string) (string, string, string) {
	company1 = fmt.Sprintf("Best Company Email Address: %s", company1)
	company2 = fmt.Sprintf("Godbless Company Email Address: %s", company2)
	company3 = fmt.Sprintf("Liberty Company Email Address: %s", company3)

	return company1, company2, company3

}
