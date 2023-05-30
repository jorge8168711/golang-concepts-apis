package main

import (
	"fmt"
	"testing"
)

// go test -> run the tests
// go test -v -> run the tests with verbose
// go test -cover -> run the tests with coverage
// go test -coverprofile=coverage.out -> run the tests with coverage and save the coverage in a file
// go tool cover -func=coverage.out -> show the coverage in the file
// go tool cover -html=coverage.out  -> show the coverage in the file in html
// go test -cpuprofile=cpu.out -> run the tests with cpu profile
// go tool pprof cpu.out
// 		top
// 		web
// 		pdf
// 		list Fibonacci

func TestSum(t *testing.T) {
	/* 	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	} */

	tables := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{20, 45, 65},
		{25, 13, 38},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.expected {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", total, item.expected)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{8, 21},
		// {50, 12586269025},
	}

	for _, item := range tables {
		result := Fibonacci(item.n)
		if result != item.expected {
			t.Errorf("Fibonacci was incorrect, got: %d, want: %d.", result, item.expected)
		}
	}
}

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  6,
			dni: "12345678A",
			mockFunc: func() {
				GetEmployById = func(id int) (Employee, error) {
					return Employee{
						Id:       6,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "Chuy",
						Age:  30,
						DNI:  "12345678A",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  30,
					DNI:  "12345678A",
					Name: "Chuy",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployById
	originalGetPersonByDNI := GetPersonByDNI

	for _, item := range table {
		item.mockFunc()

		result, err := GetFullTimeEmployeeById(item.id, item.dni)
		if err != nil {
			t.Errorf("Error getting the employee")
		}
		fmt.Println(item, result)
		if item.id != result.Id {
			t.Errorf("Error, got: %d, want: %d.", result.Id, item.id)
		}
	}

	GetEmployById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
