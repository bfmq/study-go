package main

import "fmt"

func main() {
	var s1 []int
	var s2 []string

	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	s1 = []int{1, 2, 3}
	s2 = []string{"上海"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	fmt.Printf("len:%d cap:%d\n", len(s1), cap(s1))
	fmt.Printf("len:%d cap:%d\n", len(s2), cap(s2))

	a1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	s3 := a1[0:4]
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:6]
	s6 := a1[1:]
	s7 := a1[:]
	// fmt.Println(s5)
	// fmt.Println(s6)
	// fmt.Println(s7)
	fmt.Printf("len:%d cap:%d\n", len(s5), cap(s5))
	fmt.Printf("len:%d cap:%d\n", len(s6), cap(s6))
	fmt.Printf("len:%d cap:%d\n", len(s7), cap(s7))

	s8 := s6[3:4]
	fmt.Printf("len:%d cap:%d\n", len(s8), cap(s8))

}
