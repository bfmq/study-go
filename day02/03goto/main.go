package main

import "fmt"

func main() {
bck:
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				continue bck
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
}
