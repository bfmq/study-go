package main

import "fmt"

type myInt int

func (m *myInt) hello() {
	fmt.Println("我是一个int！")
}

func main() {
	m := myInt(100)
	m.hello()
}
