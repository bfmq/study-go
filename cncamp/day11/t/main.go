package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	t := rand.Intn(2000)
	fmt.Println(t)
}
