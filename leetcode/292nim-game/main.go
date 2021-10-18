package main

import "fmt"

func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	result := canWinNim(4)
	fmt.Println(result)
	result = canWinNim(1)
	fmt.Println(result)
	result = canWinNim(2)
	fmt.Println(result)
}
