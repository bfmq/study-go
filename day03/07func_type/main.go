package main

import "fmt"

func f1() {
	fmt.Println("Hello 沙河")
}

func f2() int {
	return 10
}

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	f3(f2)
}
