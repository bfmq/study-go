package main

import (
	"fmt"
)

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("ηιδΊ")
	} else {
		fmt.Println(str)
	}

}

func main() {
	assign(100)
}
