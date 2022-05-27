package main

import "fmt"

type Receiver interface {
	Action()
}

type Receiver1 struct {
}

func (r *Receiver1) Action() {
	fmt.Println("执行命令1")
}

type Receiver2 struct {
}

func (r *Receiver2) Action() {
	fmt.Println("执行命令2")
}

type Receiver3 struct {
}

func (r *Receiver3) Action() {
	fmt.Println("执行命令3")
}

type Command interface {
	Execute()
}
type ConcreteCommand1 struct {
	receiver Receiver
}

func (c *ConcreteCommand1) Execute() {
	c.receiver.Action()
}

type ConcreteCommand2 struct {
	receiver Receiver
}

func (c *ConcreteCommand2) Execute() {
	c.receiver.Action()
}

type ConcreteCommand3 struct {
	receiver Receiver
}

func (c *ConcreteCommand3) Execute() {
	c.receiver.Action()
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) SetCommand(command Command) {
	i.commands = append(i.commands, command)
}
func (i *Invoker) ExecuteCommand() {
	for _, c := range i.commands {
		c.Execute()
	}
}

func main() {
	c1 := ConcreteCommand1{&Receiver1{}}
	c2 := ConcreteCommand1{&Receiver2{}}
	c3 := ConcreteCommand1{&Receiver3{}}

	i := Invoker{}
	i.SetCommand(&c1)
	i.SetCommand(&c2)
	i.SetCommand(&c3)
	i.ExecuteCommand()
}
