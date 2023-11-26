package main

import (
	"errors"
	"fmt"
)

const (
	color1 = "blue"
	color2 = "white"
	color3 = "black"
)

var (
	colors = map[string]string{
		color1: color1,
		color2: color2,
		color3: color3,
	}
)

type car struct {
	name  string
	color string
	model string
	year  int
}

type newCarMade struct {
	name  string
	color string
	model string
}

func newCar(n newCarMade) (*car, error) {
	car := car{name: n.name}
	car.color = n.color
	car.model = n.model
	car.year = 2023

	if _, colorIsFound := colors[car.color]; !colorIsFound {
		return nil, errors.New("no color found")
	}

	return &car, nil

}

func (c *car) GetName() string {
	return c.name

}

func (c *car) updateName(name string) *car {
	c.name = name
	return c

}

func (c *car) updateColor(color string) *car {
	c.color = color
	return c

}

func (c *car) updateModel(model string) *car {
	c.model = model
	return c
}

func main() {

	n := newCarMade{
		name:  "benz",
		color: color2,
		model: "xr",
	}

	fmt.Println(n)

	benz, err := newCar(n)
	if err != nil {
		fmt.Println(err.Error())
	}

	if benz != nil {
		fmt.Println(n)

		benz.updateName("lexus")
		benz.updateColor("color1")
		benz.updateModel("GR")
		fmt.Println(benz)
	}

}
