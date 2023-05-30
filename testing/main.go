package main

import "time"

func Sum(x, y int) int {
	return x + y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Person struct {
	Name string
	Age  int
	DNI  string
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Person
	Employee
	EndDate string
}

type TemporaryEmployee struct {
	Person
	Employee
	TaxRate int
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM person WHERE dni = dni
	return Person{}, nil
}

var GetEmployById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployById(id)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = e
	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p

	return ftEmployee, nil
}
