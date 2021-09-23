package main

import "fmt"

func fa() {
	fmt.Println("a")
}

func fb() {
	defer func() {
		// fmt.Println("释放数据库连接")
		err := recover()
		fmt.Println(err)
	}()
	panic("出现严重错误！")
	fmt.Println("b")
}

func fc() {
	fmt.Println("c")
}

func main() {
	fa()
	fb()
	fc()
}
