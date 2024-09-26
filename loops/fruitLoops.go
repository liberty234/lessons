package main

import "fmt"

func GetFruits() {

	fruit := []string{"Mango", "Orange", "Pawpaw", "Grape", "Cashew", "Pineapple", "Water melon", "Apple", "Plum", "banana", "Cherry", "strawBerry", "Papaya", "Lemon", "BlueBerry", "Kiwi", "Pear", "GrapeFruits", "Guava", "Fig", "MulBerry", "BlackBerry", "Avocado", "Mandarin", "Lime", "Melon", "Peach", "Coconut", "JackFruit", "Carambola", "Respberry", "Apricot", "Gragonfruit", "Elderberry", "Honeydew", "Lychee", "Starfruit", "Carrot", "Currant", "Key lime", "Longan", "Loquat", "Ugli fruit", "Durian", "Tangerin", "Clementin", "Quince", "Cucumber", "Squash", "Araza", "Acelora"}

	for k, v := range fruit {
		fmt.Println(k, v)

	}

}
