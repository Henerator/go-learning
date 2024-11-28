package main

import (
	"fmt"
	"golearning/19-interfaces/company"
	"golearning/19-interfaces/person"
	"golearning/19-interfaces/robot"
)

func main() {
	pers := person.Person{Name: "Jordan"}
	rob := &robot.Robot{Model: "C-3PO"}
	comp := company.Company{}

	comp.Hire(pers)
	comp.Hire(rob)

	fmt.Println(comp.Process(0, []string{"task1", "task2"}))
	fmt.Println(comp.Process(1, []string{"task3", "task4"}))
}
