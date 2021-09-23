package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("kind:%v\n", v.Kind())
	fmt.Printf("name:%v\n", v.Name())
}

func main() {
	// a := float32(1.23)
	// b := float64(4.56)
	// reflectType(a)
	// reflectType(b)
	var c = Cat{}
	reflectType(c)
}
