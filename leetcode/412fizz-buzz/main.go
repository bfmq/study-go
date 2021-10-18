package main

import (
	"fmt"
)

func fizzBuzz(n int) (result []string) {
	for i := 1; i <= n; i++ {
		var str string
		switch {
		case i%3 == 0 && i%5 == 0:
			str = "FizzBuzz"
		case i%3 == 0:
			str = "Fizz"
		case i%5 == 0:
			str = "Buzz"
		default:
			str = fmt.Sprintf("%d", i)
		}
		result = append(result, str)
	}
	return
}

func main() {
	result := fizzBuzz(3)
	fmt.Println(result)
	result = fizzBuzz(5)
	fmt.Println(result)
	result = fizzBuzz(15)
	fmt.Println(result)
}
