package main

import "fmt"

type Memento struct {
	value int
}

func NewMemento(value int) *Memento {
	return &Memento{value}
}

type Number struct {
	value int
}

func NewNumber(value int) *Number {
	return &Number{value}
}

func (n *Number) Double() {
	n.value = 2 * n.value
}
func (n *Number) Half() {
	n.value = n.value / 2
}
func (n *Number) Value() int {
	return n.value
}
func (n *Number) CreateMemento() *Memento {
	return NewMemento(n.value)
}
func (n *Number) ReinstateMemento(memento *Memento) {
	n.value = memento.value
}

func main() {
	n := NewNumber(7)
	n.Double()
	n.Double()
	memento := n.CreateMemento()
	n.Half()
	n.ReinstateMemento(memento)
	fmt.Println(n.value)
}
