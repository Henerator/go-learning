package person

import "fmt"

// Person parent structure
type Person struct {
	Name string
	Year int
}

func NewPerson(name string, year int) Person {
	return Person{
		Name: name,
		Year: year,
	}
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Birth year: %d", p.Name, p.Year)
}

func (p Person) Print() {
	fmt.Println(p)
}

// Student child structure
type Student struct {
	Person
	Group string
}

func NewStudent(name string, year int, group string) Student {
	return Student{
		Person: NewPerson(name, year),
		Group:  group,
	}
}

func (st Student) String() string {
	return fmt.Sprintf("%s, Group: %s", st.Person, st.Group)
}

func Run() {
	st := NewStudent("Jordan", 1990, "777")
	st.Print()
}
