package timeop

import "fmt"

func GetCarColor(car1, car2, car3 string) (string, string, string) {
	firstCar := fmt.Sprintf("Car:%s", car1)
	secondCar := fmt.Sprintf("Car:%s", car2)
	thirdCar := fmt.Sprintf("Car:%s", car3)
	return firstCar, secondCar, thirdCar

}
