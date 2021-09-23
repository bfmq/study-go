package main

import "fmt"

type dog struct {
	name string
}

func newDog(name string) *dog {
	return &dog{
		name: name,
	}
}

func (d *dog) wang() {
	fmt.Printf("%s:汪汪汪~\n", d.name)
}

func main() {
	d1 := newDog("金毛")
	d1.wang()
	fmt.Println(d1)
}
