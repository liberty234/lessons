package main

import "fmt"

func GetCities() {

	worldCities := []string{"Benin", "Abuja", "Lagos", "Calabar", "Ikeja", "Ibadan", "Kaduna", "Oweri", "Jos", "Sokoto", "London", "Manchester", "Fulham", "Liverpool", "Bristol", "Nottingham", "leicester", "Plymouth", "York", "Leeds", "Sheffield", "Portsmouth", "Birmigham", "Derby", "Ripon", "Ely", "Salford", "Cardiff", " St Albans", "Belfast", "Cambridge", "Southampton", "Exeter", "Norwich", "Colchester", "Wolverhampton", "Bath", "Glasgow", "Reading", "Oxford", "Chester", "Luton", "Hereford", "Boston", "Chicago", "Austin", "Portland", "New Orleans", "San Deigo", "New York", "Dallas", "Miami", "Atlanta", "Houston", "Madison", "Phoenix", "San Antonio", "Las vegas", "Omaha", "San Jose", "Savannah", "Durham", "Columbus", "Virginia Beach", "Tulsa", "Mesa", "berlin", "Bonn", "Munich", "Farnkfurt", "Cologne", "Leipzig", "Honaver", "Mainz", "Stuttgart", "Essen", "Augsburg", "Jena", "Dortmund", "Potsdom", "Paris", "Nice", "Lille", "Rennes", "Marseille", "Lyon", "cannes", "Reims", "Colmar", "Strasbourg", "Nantes", "Dijon", "Orleans", "Tours", "Arles", "Caen", "Toulouse", "Wuham", "Hefei", "Beijing", "Nanjing"}

	for K, v := range worldCities {
		fmt.Println(K, v)

	}

}
