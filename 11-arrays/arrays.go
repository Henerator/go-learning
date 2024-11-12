package main

import (
	"fmt"
)

func main() {
	createArray()
	autoLength()
	loop()
	slice()
	sliceAppend()
	task()
}

func createArray() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	arr2 := [3]int{0: 3, 2: 1}
	fmt.Println(arr2)
}

func autoLength() {
	arr := [...]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(len(arr))
}

func loop() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	var sum int = 0

	// range makes a copy of array
	for _, value := range arr {
		sum += value
	}

	fmt.Println(sum / len(arr))

	// use pointers

	var mul int = 1
	for _, value := range &arr {
		mul *= value
	}

	fmt.Println(mul)
}

func slice() {
	sl := make([]int, 5, 10)
	fmt.Println(sl)

	sl2 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl2)

	fullSlice := sl2[:]
	partSlice := sl2[1:4]
	fmt.Println("full slice", fullSlice)
	fmt.Println("part slice", partSlice)

	sl2[2] = 100
	fmt.Println("full slice", fullSlice)
	fmt.Println("part slice", partSlice)
}

func sliceAppend() {
	a := []int{1, 2, 3, 4}
	b := a[2:3]                    // b = [3]
	b = append(b, 7)               // create copy of slice b
	fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	fmt.Println(b, len(b), cap(b)) // [3 7] 2 2
	b = append(b, 8, 9, 10)
	b[0] = 11
	fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	fmt.Println(b, len(b), cap(b)) // [11 7 8 9 10] 5 6

	s := []int{1, 2, 3}
	fmt.Println(s)
	s = append(s, 4)
	fmt.Println(s)
}

func task() {
	size := 30
	sl := make([]int, size)
	for i := 0; i < size; i++ {
		sl[i] = i + 1
	}

	fmt.Println(sl)

	sl = append(sl[:10], sl[len(sl)-10-1:]...)
	fmt.Println(sl)

	size = len(sl)
	for i := range sl[:size/2] {
		sl[i], sl[size-i-1] = sl[size-i-1], sl[i]
	}
	fmt.Println(sl)
}
