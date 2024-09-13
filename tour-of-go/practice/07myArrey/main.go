package main

import "fmt"

func main() {
	fmt.Println("welcome arrey toturial in golang")

	var fruitList [4]string
	fruitList[1] = "Grape"
	fruitList[3] = "Orange"
	fruitList[2] = "pineapple"
	fruitList[0] = "mango"

	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Fruit List:", len(fruitList))

	var foodList = [6]string{"Beans,Rice,Yam,Bread"}
	fmt.Println("Food List:", foodList)
	fmt.Println("Food List:", len(foodList))
	fmt.Printf("type of data %T \n", foodList)

	var name = [3]string{"Liberty", "John", "Fred"}
	fmt.Println("List of Name:", name)
	fmt.Println("list of name:", len(name))
	fmt.Printf("type of data %T", name)

}
