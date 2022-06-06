package main

import "fmt"

// 鼠标接口
type mouse interface {
	click()
}
type AsusMouse struct {
}

func (a *AsusMouse) click() {
	fmt.Println("Asus Mouse Check screen")
}

type LogitechMouse struct {
}

func (l *LogitechMouse) click() {
	fmt.Println("Logitech Mouse Check screen")
}

// 键盘接口
type keyboard interface {
	press()
}
type AsusKeyboard struct {
}

func (a *AsusKeyboard) press() {
	fmt.Println("Asus keyboard press")
}

type LogitechKeyboard struct {
}

func (l *LogitechKeyboard) press() {
	fmt.Println("Logitech Keyboard press")
}

// 抽象工厂，抽象了鼠标/键盘生成的工厂
type AbstractFactory interface {
	produceMouse(brand string) mouse
	produceKeyboard(brand string) keyboard
}

// 鼠标工厂
type MouseFactory struct {
}

func (m *MouseFactory) produceMouse(brand string) mouse {
	switch brand {
	case "Asus":
		return &AsusMouse{}
	case "Logitech":
		return &LogitechMouse{}
	}
	return nil
}

// 鼠标工厂无法生成键盘
func (m *MouseFactory) produceKeyboard(brand string) keyboard {
	return nil
}

// 键盘工厂
type KeyboardFactory struct {
}

func (k KeyboardFactory) produceMouse(brand string) mouse {
	return nil
}

func (k KeyboardFactory) produceKeyboard(brand string) keyboard {
	switch brand {
	case "Asus":
		return &AsusKeyboard{}
	case "Logitech":
		return &LogitechKeyboard{}
	}
	return nil
}

// 抽象工厂的简单工厂
func NewFactory(factory string) AbstractFactory {
	switch factory {
	case "mouse":
		return &MouseFactory{}
	case "keyboard":
		return &KeyboardFactory{}
	}
	return nil
}

func main() {
	factory := NewFactory("mouse")
	asusMouse := factory.produceMouse("Asus")
	asusMouse.click()
}
