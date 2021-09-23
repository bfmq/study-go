package main

import (
	"fmt"
)

func main() {

	s := "Hello沙河"
	n := len(s)
	fmt.Println(n)

	// for i := 0; i < n; i++ {
	// 	fmt.Println(s[i])
	// 	// fmt.Printf("%c", s[i])
	// }

	// for _, c := range s {
	// 	// fmt.Println(i)
	// 	// fmt.Println(c)
	// 	fmt.Printf("%c", c)
	// }

	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))

	n1 := 10
	f := float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
