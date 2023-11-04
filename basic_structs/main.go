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

		fmt.Println("Car Name:", lexus.name)
		fmt.Println("Car Color:", lexus.color)
		fmt.Println("Car Model:", lexus.model)
		fmt.Println()
		lexus = lexus.UpdateName("Lexus360")
		lexus = lexus.UpdateColor(color2)
		lexus = lexus.UpdateModel("RC")
		lexus = lexus.UpdateYear(2024)
		fmt.Println(lexus)

	}

	fmt.Println()

	employ := secondEmployee{
		name:   "Mike",
		age:    28,
		job:    "Software Engineer",
		salary: "200k",
	}

	fmt.Println(employ)

	employee, err := newEmployee(employ)

	if err != nil {
		fmt.Println(err.Error())
	}

	if employee != nil {
		fmt.Println("Employee Name:", employee.name)
		fmt.Println("Employee Age:", employee.age)
		fmt.Println("Employee Job:", employee.job)
		fmt.Println("Employee Salary:", employee.salary)
		fmt.Println("The new Update")
		employee = employee.UpdateName("Sam")
		employee = employee.UpdateAge(25)
		employee = employee.UpdateJob("Banker")
		employee = employee.UpdateSalary("100k")
		fmt.Println(employee)

	}

}
