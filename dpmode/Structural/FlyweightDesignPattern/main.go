package main

import "fmt"

type FlyWeight struct {
	Name string
}

func NewFlyWeight(name string) *FlyWeight {
	return &FlyWeight{Name: name}
}

type FlyWeightFactory struct {
	pool map[string]*FlyWeight
}

func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{pool: make(map[string]*FlyWeight)}
}

func (f *FlyWeightFactory) GetFlyWeight(name string) *FlyWeight {
	weight, ok := f.pool[name]
	if !ok {
		weight = NewFlyWeight(name)
		f.pool[name] = weight
	}
	return weight
}

func main() {
	factory := NewFlyWeightFactory()
	hong := factory.GetFlyWeight("hong beauty")
	xiang := factory.GetFlyWeight("xiang beauty")
	fmt.Println(len(factory.pool))
	fmt.Println(hong)
	fmt.Println(xiang)
}
