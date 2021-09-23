package main

import "fmt"

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿")
	fmt.Println("end")
}

func main() {
	deferDemo()
}
