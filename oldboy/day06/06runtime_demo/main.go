package main

import (
	"fmt"
	"runtime"
)

func f() {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Printf("runtime.Caller() failed,err:\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(line)

}

func f1() {
	f()
}
func main() {
	f1()
}
