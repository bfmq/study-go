package main

import "fmt"

type animal struct {
	name string
}

type dog struct{
	feet uint8
	animal
}

func (a *animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

func (d *dog)wang()  {
	fmt.Printf("%s汪汪汪\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "金毛",
		},
	}
	d1.wang()
	d1.move()
}
