package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Println("上线")
	} else {
		fmt.Println("暑假作业")
	}

	if age > 35 {
		fmt.Println("中年")
	} else if age > 18 {
		fmt.Println("青年")
	}else {
		fmt.Println("暑假作业")
	}
}
