package main

import "fmt"

func main() {
	threeComponents()
	oneCondition()
	rangeLoop()
	labels()
}

func threeComponents() {
	var value = 0
	for i := 0; i < 10; i++ {
		value++
	}

	println(value)
}

func oneCondition() {
	var i = 0
	for i < 5 {
		i++
	}
	println(i)
}

func rangeLoop() {
	var array = [3]int{1, 2, 3}
	for index, value := range array {
		fmt.Printf("array[%d]: %d\n", index, value)
	}
}

func labels() {
outerLabel:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			break outerLabel
		}
	}
}
