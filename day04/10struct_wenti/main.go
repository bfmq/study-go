package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("a", 18)
	p2 := newPerson("撒", 1)
	fmt.Println(p1, p2)
}
