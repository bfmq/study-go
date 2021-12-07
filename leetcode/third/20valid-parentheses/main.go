package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	for strings.Index(s, "()") != -1 || strings.Index(s, "[]") != -1 || strings.Index(s, "{}") != -1 {
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")
	}
	return len(s) == 0
}

func main() {
	result := isValid("[][][]")
	fmt.Println(result)
	result = isValid("()[]{}")
	fmt.Println(result)
	result = isValid("(]")
	fmt.Println(result)
	result = isValid("([)]")
	fmt.Println(result)
	result = isValid("{[]}")
	fmt.Println(result)
}
