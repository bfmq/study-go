package main

import "fmt"

type WorkInterface interface {
	GetUp()
	Work()
	Sleep()
}
type Worker struct {
	WorkInterface
}

func NewWorker(w WorkInterface) *Worker {
	return &Worker{WorkInterface: w}
}

// 模版方法
func (w *Worker) Daily() {
	w.GetUp()
	w.Work()
	w.Sleep()
}

// 实现结构体
type Coder struct{}

// 具体每步实际工作
func (c *Coder) GetUp() {
	fmt.Println("Coder GetUp ")
}

func (c *Coder) Work() {
	fmt.Println("Coder Work ")
}

func (c *Coder) Sleep() {
	fmt.Println("Coder Sleep ")
}

func main() {
	c := new(Coder)
	w := NewWorker(c)
	w.Daily()
}
