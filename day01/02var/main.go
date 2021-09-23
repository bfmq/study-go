package main

import "fmt"

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "bfmq"
	age = 29
	isOk = true
	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name: %s\n", name)
	fmt.Println(age)

	s1 := "game"
	fmt.Println(s1)

}
