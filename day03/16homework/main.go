package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatch(userName string) int {
	userCoin := 0
	if strings.Count(userName, "e") == 1 || strings.Count(userName, "E") == 1 {
		userCoin += 1
	}
	if strings.Count(userName, "i") == 1 || strings.Count(userName, "I") == 1 {
		userCoin += 2
	}
	if strings.Count(userName, "o") == 1 || strings.Count(userName, "O") == 1 {
		userCoin += 3
	}
	if strings.Count(userName, "u") == 1 || strings.Count(userName, "U") == 1 {
		userCoin += 4
	}

	return userCoin
}

func dispatchCoin() int {
	for _, userName := range users {
		userCoin := dispatch(userName)
		// println(userName, userCoin)
		distribution[userName] = userCoin
		coins -= userCoin
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)
}
