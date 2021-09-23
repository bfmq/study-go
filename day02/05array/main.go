package main

import "fmt"

func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)
	fmt.Println(a1)
	fmt.Println(a2)

	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	a10 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a10)

	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	citys := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	fmt.Println("===================================")
	for _, v := range citys {
		fmt.Println(v)
	}

	fmt.Println("===================================")
	a11 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)
	for _, v1 := range a11 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1)
	fmt.Println(b2)

}
