package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "10000"
	retInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf(" err:%s\n", err)
		return
	}
	fmt.Println(retInt)

	boolStr := "true"
	retBool, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Printf(" err:%s\n", err)
		return
	}
	fmt.Println(retBool)
}
