package main

import "fmt"

type AbstractDog interface {
	GetName() string
	GetPrice() float64
}

//具体构件
type BorderCollie struct{}

func (b *BorderCollie) GetName() string {
	return "边牧"
}

func (b *BorderCollie) GetPrice() float64 {
	return 2600
}

//抽象装饰角色
type DogDecorator struct {
	dog AbstractDog
}

//具体装饰A
type DogFood struct {
	dogDecorator DogDecorator
}

func (f *DogFood) GetName() string {
	return f.dogDecorator.dog.GetName() + "狗粮"
}

func (f *DogFood) GetPrice() float64 {
	return f.dogDecorator.dog.GetPrice() + 100
}

func main() {
	bc := new(BorderCollie)
	dogfood := new(DogFood)
	dogfood.dogDecorator.dog = bc
	fmt.Printf("%v, %v\n", dogfood.GetName(), dogfood.GetPrice())
}
