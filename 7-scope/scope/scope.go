package main

import (
	"fmt"
	"golearning/7-scope/lib"
)

func main() {
	lib.SetSupport("Служба поддержки")
	fmt.Println(lib.GetContact())
	fmt.Println("Email:", lib.Email)
}
