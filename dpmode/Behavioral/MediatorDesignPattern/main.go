package main

import "fmt"

type Colleague interface {
	Send(msg string)
	Notify(msg string)
	SetMediator(mediator Mediator)
}
type ConcreteColleague1 struct {
	mediator Mediator
}

func (c *ConcreteColleague1) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ConcreteColleague1) Send(msg string) {
	c.mediator.Send(msg, c)
}

func (c *ConcreteColleague1) Notify(msg string) {
	fmt.Println("ConcreteColleague1 recv msg:", msg)
}

type ConcreteColleague2 struct {
	mediator Mediator
}

func (c *ConcreteColleague2) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ConcreteColleague2) Send(msg string) {
	c.mediator.Send(msg, c)
}

func (c *ConcreteColleague2) Notify(msg string) {
	fmt.Println("ConcreteColleague2 recv msg:", msg)
}

type Mediator interface {
	Send(msg string, colleague Colleague)
}
type ConcreteMediator struct {
	c1 Colleague
	c2 Colleague
}

func (c *ConcreteMediator) Send(msg string, colleague Colleague) {
	if colleague == c.c1 {
		c.c2.Notify(msg)
	} else {
		c.c1.Notify(msg)
	}
}

func main() {
	concreteMediator := &ConcreteMediator{}
	c1 := &ConcreteColleague1{}
	c2 := &ConcreteColleague2{}
	c1.SetMediator(concreteMediator)
	c2.SetMediator(concreteMediator)
	concreteMediator.c1 = c1
	concreteMediator.c2 = c2
	c1.Send("吃饭了吗c2")
	c2.Send("我吃过了c1")
}
