package main

import (
	"fmt"
	"golearning/20-generator/randbyte"
	"time"
)

func main() {
	buf := make([]byte, 16)
	generator := randbyte.NewGenerator(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		n, _ := generator.Read(buf)
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}
}
