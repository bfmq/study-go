package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	// s1[3] = "广州"
	fmt.Printf("s1 len:%d,s1 cap:%d\n", len(s1), cap(s1))

	s1 = append(s1, "广州")
	fmt.Printf("s1 len:%d,s1 cap:%d\n", len(s1), cap(s1))

}
