package main

import (
	"fmt"
	"time"
)

func main() {
	casefall()
	greeting(1960)
	greeting(2024)
	fmt.Println(getYear())
}

func casefall() {
	a := -100
	switch {
	case a > 0:
		if a%2 == 0 {
			break
		}
		fmt.Println("Odd positive value received")
	case a < 0:
		fmt.Println("Negative value received")
		fallthrough
	default:
		fmt.Println("Default value handling")
	}
}

func greeting(bornYear int) {
	switch {
	case bornYear <= 1964:
		fmt.Println("Bumer")
	case bornYear <= 1980:
		fmt.Println("X")
	case bornYear <= 1996:
		fmt.Println("Millenial")
	case bornYear <= 2012:
		fmt.Println("Zoomer")
	default:
		fmt.Println("Alpha")
	}
}

func getYear() int {
	var date = time.Date(2000, 10, 1, 0, 0, 0, 0, time.UTC)
	return date.Year()
}
