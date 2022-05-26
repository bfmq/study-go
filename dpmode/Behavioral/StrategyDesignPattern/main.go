package Strategy

import "fmt"

type Strategy interface {
	Execute()
}

type strategyA struct{}

func NewStrategyA() Strategy {
	return &strategyA{}
}

func (s *strategyA) Execute() {
	fmt.Println("A plan executed.")
}

type strategyB struct{}

func (s *strategyB) Execute() {
	fmt.Println("B plan executed.")
}

func NewStrategyB() Strategy {
	return &strategyB{}
}

type Context struct {
	strategy Strategy
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) Execute() {
	c.strategy.Execute()
}

func main() {
	c := NewContext()
	strategyA := NewStrategyA()
	strategyB := NewStrategyB()

	c.SetStrategy(strategyA)
	c.Execute()

	c.SetStrategy(strategyB)
	c.Execute()
}
