package main

import "fmt"

const (
	statusOK = 200
	notFound = 404
)

const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota
	a2
	_
	a3
)

func main() {
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

}
