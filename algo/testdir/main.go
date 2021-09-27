package main

import "fmt"

func main() {
	data := make([]interface{}, 0, 32)
	fmt.Println(len(data))
	fmt.Println(cap(data))
}
