package main

import "fmt"

type Target interface {
	Execute()
}

type Source struct{}

func (a *Source) SepcificExecute() {
	fmt.Println("原充电...")
}

type Adapter struct {
	*Source
}

func (a *Adapter) Execute() {
	a.SepcificExecute()
}

func main() {
	a := Adapter{}
	a.Execute()
}
