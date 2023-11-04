package main

import "fmt"

func FetchLanguage() {

	language := make(map[string]string)
	language["JS"] = "Javascript"
	language["RB"] = "Ruby"
	language["Py"] = "Python"

	fmt.Println("List all languages:", language)
	delete(language, "RB")
	fmt.Println("List all languages:", language)

}
