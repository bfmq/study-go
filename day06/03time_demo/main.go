package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// fmt.Println(now.Year())
	// fmt.Println(now.Month())
	// fmt.Println(now.Date())
	// fmt.Println(now.Hour())
	// fmt.Println(now.Minute())
	// fmt.Println(now.Second())
	// fmt.Println(now.Unix())
	// fmt.Println(now.UnixNano())
	// fmt.Println(now.Add(4 * time.Hour))
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	fmt.Println(now.Format("2006-01-02"))
}
