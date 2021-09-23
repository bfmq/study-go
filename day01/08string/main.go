package main

import (
	"fmt"
	"strings"
)

func main() {

	path := "'C:\\Users\\Aone\\go\\src'"
	// fmt.Println(path)

	fmt.Println(len(path))

	name := "asdafsad"
	world := "dasdasdasss"

	ss := name + world
	fmt.Println(ss)

	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)

	ret := strings.Split(path, "\\")
	fmt.Println(ret)

	fmt.Println(strings.Contains(ss, "s"))
	fmt.Println(strings.Contains(ss, "ssssssssssss"))
	
	fmt.Println(strings.HasPrefix(ss, "as"))
	fmt.Println(strings.HasSuffix(ss, "as"))

	s4:="abcdeb"
	fmt.Println(strings.Index(s4, "b"))
	fmt.Println(strings.LastIndex(s4, "b"))

	fmt.Println(strings.Join(ret, "+"))

}
