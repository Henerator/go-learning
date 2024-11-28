package person

import "fmt"

type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

func (p Person) Work(tasks []string) string {
	s := p.Name + " work:"
	for _, task := range tasks {
		s += "\n I do " + task
	}

	return s
}

func Run() {
	p := Person{Name: "Jordan"}
	fmt.Println(p)
}
