package main

import "fmt"

type Prototype struct {
	Elem    string
	Content string
}

func (p *Prototype) Copy() *Prototype {
	return &(*p)
}

func main() {
	ptObj := new(Prototype)
	ptObj.Elem = "this is prototype"
	ptCopy := ptObj.Copy()
	ptCopy.Content = "this is prototype clone "

	fmt.Println(ptObj)
	fmt.Println(ptCopy)
}
