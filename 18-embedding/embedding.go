package main

import (
	"fmt"
	"golearning/18-embedding/logger"
	"golearning/18-embedding/person"
)

func main() {
	fmt.Println("[Person Section]")
	person.Run()

	logger := logger.NewExtendedLog()
	logger.SetLogLevel(1)
	logger.Infoln("Test info")
	logger.Warnln("Test warning")
	logger.Errorln("Test error")
}
