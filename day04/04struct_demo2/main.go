package main

import "fmt"

type person struct {
	name, gender string
}

func f(x person) {
	x.gender = "女"
}

func f2(x *person) {
	(*x).gender = "女"
}

func main() {
	var p person
	p.name = "周大"
	p.gender = "男"
	f(p)
	fmt.Println(p)
	f2(&p)
	fmt.Println(p)

	p3 := person{
		name:   "把娃",
		gender: "女",
	}
	fmt.Println(p3)
}
