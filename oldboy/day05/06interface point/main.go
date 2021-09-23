package main

import (
	"fmt"
)

type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(something string) {
	fmt.Printf("猫吃%s\n", something)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动")
}

func (c chicken) eat(something string) {
	fmt.Println("鸡吃饲料")
}

func main() {
	var a1 animal
	bc := cat{
		name: "淘气",
		feet: 4,
	}

	kfc := chicken{
		feet: 2,
	}

	a1 = bc
	a1.eat("小鱼仔")

	a1 = kfc
	a1.eat("小鱼仔")

}
