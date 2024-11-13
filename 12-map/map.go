package main

import (
	"fmt"
)

func main() {
	createMap()
	getValue()
	task1()
	task2()
	task3()
}

func createMap() {
	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"
	fmt.Println(m)

	mComposite := map[string]string{"key1": "value1"}
	fmt.Println(mComposite)
}

func getValue() {
	m := map[string]int{"key1": 1}
	noExistValue, ok := m["no-key"]
	fmt.Println(noExistValue, ok)

	delete(m, "key1")
	fmt.Println(m)
}

func task1() {
	m := make(map[string]int)
	m["bread"] = 50
	m["milk"] = 100
	m["oil"] = 200
	m["sausage"] = 500
	m["salt"] = 20
	m["cucumber"] = 200
	m["cheese"] = 600
	m["ham"] = 700
	m["tomato"] = 250
	m["fish"] = 300

	for key, value := range m {
		if value > 500 {
			fmt.Println(key, value)
		}
	}

	items := []string{"bread", "cheese", "cucumber"}
	sum := 0

	for _, item := range items {
		sum += m[item]
	}
	fmt.Println("sum:", sum)
}

func task2() {
	arr := []int{1, 2, 3, 4, 5, 6}
	k := 9
	m := map[int]int{}

	for i, value := range arr {
		rest := k - value
		if j, ok := m[rest]; ok {
			fmt.Println("i:", i, "j:", j)
			return
		}

		m[value] = i
	}

	fmt.Println("not found")
}

func task3() {
	removeDuplicate := func(arr []string) []string {
		output := make([]string, 0)
		m := make(map[string]int, len(arr))

		for _, value := range arr {
			if _, ok := m[value]; !ok {
				output = append(output, value)
				m[value] = 1
			}
		}

		return output
	}

	fmt.Println(removeDuplicate([]string{"1", "1", "2", "3", "3"}))
}
