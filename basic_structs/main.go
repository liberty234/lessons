package main

import "fmt"

func main() {

	m := newCarMade{
		name:  "Lexus350",
		color: color1,
		model: "XR",
	}

	fmt.Println(m)

	lexus, err := newCar(m)

	if err != nil {
		fmt.Println(err.Error())
	}

	if lexus != nil {
		fmt.Println(lexus)
		fmt.Println("Car Name:", lexus.name)
		fmt.Println("Car Color:", lexus.color)
		fmt.Println("Car Model:", lexus.model)
		fmt.Println()
		lexus = lexus.UpdateName("Lexus360")
		lexus = lexus.UpdateColor(color2)
		lexus = lexus.UpdateModel("RC")
		fmt.Println(lexus)

	}

}
