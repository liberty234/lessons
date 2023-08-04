package arrey

import "fmt"

func getFoodStuff() {
	foodStuff := []string{"Beans", "Rice", "Pepper", "Maggi Cube"}

	foodStuff[1] = "Garri"
	foodStuff = append(foodStuff, "Ginger", "curry")

	fmt.Println("FoodStuff:", foodStuff)

}
