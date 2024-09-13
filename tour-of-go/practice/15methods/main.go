package main

import "fmt"

type Details struct {
	Name       string
	Email      string
	Number     int
	HasGlasses bool
	Age        int
}

func main() {
	fmt.Println("Struct in Golang")

	details := Details{"Liberty", "liberty@mail.com", 7061267454, false, 18}
	fmt.Println(details)
	fmt.Printf("All details: %+v\n", details)
	details.GetName()
	fmt.Println("New Email:", details.NewEmail())

}

func (d Details) GetName() {
	fmt.Println("Name:", d.Name)

}

func (d Details) NewEmail() string {
	d.Email = "ebikadeli2020@gmail.com"
	return d.Email

}
