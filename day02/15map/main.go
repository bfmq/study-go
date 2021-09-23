package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var socreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		socreMap[key] = value
	}

	var keys = make([]string, 0, 200)
	for key := range socreMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		println(key, socreMap[key])
	}
}
