// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func main() {
	result := replaceSpace("We are happy.")
	fmt.Println(result)
}
