package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// var i = 5
	// for ;i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// var i = 5
	// for ;i < 10; {
	// 	fmt.Println(i)
	// 	i++
	// }


	s := "Hello沙河"
	for i, v := range s{
		fmt.Printf("%d%c\n", i, v)
	}

}
