package main

import "fmt"

func main() {
	s1 := make([]int,5,10)
	fmt.Printf("s1 len:%d,s1 cap:%d\n", len(s1), cap(s1))
}
