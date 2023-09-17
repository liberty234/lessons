package timeop

import "fmt"

func GetCarColor(car1, car2, car3 string) (string, string, string) {
	f := fmt.Sprintf("Car:%s", car1)
	e := fmt.Sprintf("Car:%s", car2)
	g := fmt.Sprintf("Car:%s", car3)
	return f, e, g

}
