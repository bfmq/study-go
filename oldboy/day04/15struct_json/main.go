package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// type dog struct{
// 	feet uint8
// 	animal
// }

func main() {
	p1 := person{
		Name: "北方",
		Age:  29,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%v\n", string(b))

	str := `{"name": "跑路", "age": 44}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%v\n", p2)


}
