package main

import "fmt"

type Car interface {
	drive()
}

type bmw struct {
	name string
}

func (b *bmw) drive() {
	fmt.Println("正在驾驶宝马...")
}

type tesla struct {
	name string
}

func (t *tesla) drive() {
	fmt.Println("正在驾驶特斯拉...")
}

type porsche struct {
	name string
}

func (p *porsche) drive() {
	fmt.Println("正在驾驶保时捷...")
}

func CreateCarFactory(car string) Car {
	switch car {
	case "bmw":
		return &bmw{}
	case "tesla":
		return &tesla{}
	case "porsche":
		return &porsche{}
	default:
		return nil
	}
}

func main() {
	car := CreateCarFactory("bmw")
	car.drive()
	car = CreateCarFactory("tesla")
	car.drive()
	car = CreateCarFactory("porsche")
	car.drive()
}
