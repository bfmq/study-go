package main

import "fmt"

func f1() {
	fmt.Println("i'm f1")
}

func main() {
	ch := make(chan bool, 3)
	fmt.Println(ch)
	go f1()
}
