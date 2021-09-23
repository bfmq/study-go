package main

/*
	函数版学生管理
*/

import "fmt"

type student struct {
	id   int64
	name string
}

func main() {
	fmt.Println("欢迎使用学生管理系统")
	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
	`)

	

}
