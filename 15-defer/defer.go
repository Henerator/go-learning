package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	declare()
	readFile()
	handlePanic()
}

func declare() {
	defer fmt.Println("defer")
	defer fmt.Println("defer2")
	fmt.Println("declare")
}

func readFile() {
	// открываем файл
	file, err := os.OpenFile("./15-defer/file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("could not open file", err)
		return
	}

	// не забываем закрыть файл
	defer func() {
		file.Close()
		printDuration(time.Now())
	}()

	// работаем с файлом
	_, err = file.WriteString("")
	if err != nil {
		fmt.Println("could not write to file", err)
	}
}

func printDuration(start time.Time) {
	fmt.Println("duration: ", time.Since(start))
}

func handlePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic was caught", r)
		}
	}()

	panic("Don't use panic like this")
}
