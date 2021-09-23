package main

import (
	"fmt"
)

type car interface {
	run()
}

type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

func (f falali) run() {
	fmt.Printf("%s速度70\n", f.brand)
}

func (b baoshijie) run() {
	fmt.Printf("%s速度700\n", b.brand)
}

func drive(c car) {
	c.run()
}

func main() {
	var f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}

	drive(f1)
	drive(b1)
}
