package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Mul(2, 3))
	fmt.Println(Mul("2", 3))
}

func Mul(a interface{}, b int) interface{} {
	switch value := a.(type) {
	case int:
		return value * b
	case string:
		return strings.Repeat(value, b)
	case fmt.Stringer:
		return strings.Repeat(value.String(), b)
	default:
		return nil
	}
}
