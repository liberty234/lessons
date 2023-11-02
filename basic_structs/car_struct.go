package main

import (
	"errors"
)

const (
	color1 = "Yellow"
	color2 = "Black"
	color3 = "White"
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

func newCar(m newCarMade) (*car, error) {
	car := car{name: m.name}
	car.color = m.color
	car.model = m.model
	car.year = 2023

	if _, colorIsFound := colors[car.color]; !colorIsFound {
		return nil, errors.New("please choose the right color")
	}
	return &car, nil

}

func (c *car) GetName(string) string {
	return c.name
}

func (c car) GetColor() string {
	return c.color
}

func (c car) GetModel() string {
	return c.model
}

func (c *car) UpdateName(name string) *car {

	c.name = name

	return c

}

func (c *car) UpdateColor(color string) *car {

	c.color = color
	return c

}

func (c *car) UpdateModel(model string) *car {

	c.model = model
	return c

}
