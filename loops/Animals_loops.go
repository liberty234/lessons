package main

import "fmt"

func GetAnimals() {

	animals := []string{"Fowl", "Goat", "Sheep", "Rabbir", "Rat", "Elephant", "Turkey", "Whale", "Cricket", "Cow", "Pig", "Parrot", "Tiger", "Lion", "Bear", "Ape", "Cat", "Wolf", "Monkey", "Dear", "Hippopotamus", "Giraffe", "Girilla", "chipazin", "Fpx", "Loepard", "Kangaroo", "Horse", "Koala", "Camel", "Donkey", "Bat", "Snake", "Crocodile", "Sqirrel", "Peacock", "beetle", "Butterfly", "Dove", "Rhino", "Buffalo", "Hen", "Cock", "Ostrich", "Antelope", "Frog", "Eagle", "Pigeon", "Zebra", "Duck"}

	for k, v := range animals {
		fmt.Println(k, v)

	}

}
