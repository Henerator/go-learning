package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	assignPointer()
	composite()
	inc()
	person()
	inputCount()
}

func assignPointer() {
	var a int = 1
	var p *int = &a

	println(p)  //0xc000057f30
	println(*p) //1
}

func composite() {
	type A struct {
		IntField int
	}

	p := &A{IntField: 2}

	// p := new(A)
	// p := &A{}

	println(p)          //0xc000057f38
	println(p.IntField) // =(*p).IntField
}

func inc() {
	incrementCopy := func(i int) {
		i++
	}

	increment := func(i *int) {
		(*i)++
	}

	i := 42

	incrementCopy(i)
	println(i) // 42

	increment(&i)
	println(i) // 43
}

func person() {
	type User struct {
		Name       string
		lastChange time.Time
	}

	udpateLastChange := func(user *User) {
		user.lastChange = time.Now()
	}

	user := User{
		Name: "name",
	}
	udpateLastChange(&user)
	fmt.Println(user.lastChange)
}

func inputCount() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	f := func(count *int) {
		(*count)++
	}

	cnt := 0
	for {
		fmt.Print("-> ")
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}
}
