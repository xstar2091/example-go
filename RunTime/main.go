package main

import (
	"fmt"
	"path"
	"runtime"
)

func log(msg string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Could not get caller information")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file)
	fmt.Printf("[%s:%d] %s: %s\n", fileName, line, funcName, msg)
}

func main() {
	log("This is a test message")
	log("222")
	log("333")
	log("444")
	log("555")
}
