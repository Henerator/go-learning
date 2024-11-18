package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	factorial()
	fibbonachi()
	PrintAllFiles("./assets", func(filename string) bool {
		return strings.Contains(filename, "co")
	})
	generator()
}

func getFactorial(n int) int {
	if n == 0 {
		return 1
	}

	return getFactorial(n-1) * n
}

func factorial() {
	fmt.Println(getFactorial(4))
}

func getFibbonachi(n int) int {
	a, b := 0, 1
	for n > 0 {
		a, b = b, a+b
		n--
	}

	return a
}

func fibbonachi() {
	fmt.Println(getFibbonachi(7))
}

func PrintAllFiles(path string, filter func(string) bool) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())
		// печатаем имя элемента
		if filter(filename) {
			fmt.Println(filename)
		}
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFiles(filename, filter)
		}
	}
}

func getGenerator(start, step int) func() int {
	return func() int {
		value := start
		start += step
		return value
	}
}

func generator() {
	generator := getGenerator(1, 2)
	fmt.Println("Generator:")
	fmt.Println(generator())
	fmt.Println(generator())
	fmt.Println(generator())
}
