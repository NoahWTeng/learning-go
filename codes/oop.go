package main

import (
	"fmt"
	"strconv"
)

type Employee struct {
	name string
	age int
	phone string
	role string
}
type Manager struct {
	Employee
	salary int
}

type Programmer struct {
	Employee
	salary int
	workingHour int
}

type Salesman struct {
	Employee
	salary int
	sales int
}

func (v Manager) whoIam()  {
	fmt.Printf("Hi, I am %s, I am a %s of Goland. Call me on %s\n", v.name,
		v.role, v.phone)
}

func (v Salesman) whoIam()  {
	fmt.Printf("Hi, I am %s, I am a %s of Goland. Call me on %s\n", v.name,
		v.role, v.phone)
}

func (v Programmer) whoIam()  {
	fmt.Printf("Hi, I am %s, I am a %s of Goland. Call me on %s\n", v.name,
		v.role, v.phone)
}

func (v Manager) getTotalSalary()  {
	fmt.Printf("My total salary is: %s \n", strconv.Itoa(v.salary))
}

func (v Salesman) getTotalSalary()  {
	fmt.Printf("My total salary is: %s  \n", strconv.Itoa(v.salary * v.sales))
}

func (v Programmer) getTotalSalary()  {
	fmt.Printf("My total salary is: %s \n", strconv.Itoa(v.salary * v.workingHour))
}


type Info interface {
	whoIam()
	getTotalSalary()
}

func main()  {
	noah := Manager{Employee{"Noah", 33, "333-444-666", "Manager"}, 25000}
	anna := Salesman{Employee{"Anna", 26, "222-666-777", "Saleswoman"}, 500, 10}
	josh := Programmer{Employee{"Josh", 28, "111-666-888", "Programmer"}, 1000, 12}

	var _ Info
	_ = noah
	noah.whoIam()
	noah.getTotalSalary()


	_ = anna
	anna.whoIam()
	anna.getTotalSalary()

	_ = josh
	josh.whoIam()
	josh.getTotalSalary()


}

