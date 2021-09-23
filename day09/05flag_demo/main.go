package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "备份", "请输入名字")
	age := flag.Int("age", 300,"请输入年龄")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
}
