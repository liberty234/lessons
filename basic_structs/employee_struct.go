package main

type employee struct {
	name   string
	age    int
	job    string
	salary string
}

type secondEmployee struct {
	name   string
	age    int
	job    string
	salary string
}

func newEmployee(employ secondEmployee) (*employee, error) {
	employee := employee{name: employ.name}
	employee.age = employ.age
	employee.job = employ.job
	employee.salary = employ.salary

	return &employee, nil

}

func (e employee) GetName() string {
	return e.name

}

func (e employee) GetAge() int {
	return e.age

}

func (e employee) GetJob() string {
	return e.job

}

func (e employee) GetSalary() string {
	return e.salary

}

func (e *employee) UpdateName(name string) *employee {
	e.name = name
	return e

}

func (e *employee) UpdateAge(age int) *employee {
	e.age = age
	return e

}

func (e *employee) UpdateJob(job string) *employee {
	e.job = job
	return e

}

func (e *employee) UpdateSalary(salary string) *employee {
	e.salary = salary
	return e

}
