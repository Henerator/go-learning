package main

type MyString = string // alias for string
type TestString string // custom type

type Weekday int

func main() {
	var (
		str        = "Hellow World"
		ver string = "v0.0.1"
	)

	id := 0
	const pi = 3.1415

	println(string(str[0]))
	println("ver =", ver, "id =", id, "pi =", pi)

	constants()
	enum()
	weekday()
}

func constants() {
	const id = 100     // constant is untyped
	var i int64 = id   // and can be assigned
	var f float64 = id // to different types

	println("i=", i, "f=", f)
}

func enum() {
	const (
		A = iota
		B
		C
	)

	const (
		_ = iota * 10
		ten
		twenty
		thirty
	)

	const (
		hello = "Hello, world!" // iota равна 0
		one   = 1               // iota равна 1
		black = iota            // iota равна 2
		gray
	)

	println(A, B, C)
	println(ten, twenty, thirty)
	println(hello, one, black, gray)
}

func NextDay(day Weekday) Weekday {
	return (day % 7) + 1
}

func weekday() {
	const (
		Monday Weekday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	var today Weekday = Sunday
	var tomorrow = NextDay(today)
	println("today =", today, "tomorrow =", tomorrow)
}
