package main

import "fmt"

// var f1 = func(x, y int) {
// 	fmt.Println(x + y)
// }

func main() {
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)
}
