package main

import "fmt"

func main() {
	// var s1 = make([]map[int]string, 10)
	// s1[0] = make(map[int]string, 1)
	// s1[0][100] = "A"
	// s1[0][200] = "A"
	// fmt.Println(s1)

	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)
}
