package main

import (
	"fmt"
	"reflect"
)

type People struct {
	x string
	y string
}	

func (p *People) F1() {
	fmt.Println("f1")
}

func (p *People) F2() {
	fmt.Println("f2")
}

func main() {
	p := &People{
		x: "x",
		y: "y",
	}

	v := reflect.ValueOf(*p)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("field %d is %s\n", i, v.Field(i))
	}
	fmt.Println("==========")

	t := reflect.TypeOf(p)
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("method %d is %s\n", i, t.Method(i).Name)
	}
}
