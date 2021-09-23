package main

import (
	"fmt"
)

func main() {
	m1 := make(map[string]int, 10)
	m1["理想"] = 18
	m1["jiwuming"] = 35
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	fmt.Println(m1["娜扎"])
	v, ok := m1["娜扎"]
	if !ok{
		fmt.Println("查无此人！")
	}else{
		fmt.Println(v)
	}
}
