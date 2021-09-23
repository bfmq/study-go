package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:nrXTTQguLpXJGIkrExklwecwiQGQsjqYM1OVM35YHf12aK6LE5EnlopeWrno@tcp(10.10.2.20:3306)/go_test"
	db, err := sql.Open("mysql", dsn)
	if err != nil { // 只是验证dsn格式
		fmt.Printf("dsn:%s invalid,err:%s\n", dsn, err)
		return
	}

	err = db.Ping() // 尝试链接数据库
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%s\n", dsn, err)
		return
	}
	defer db.Close()
	fmt.Println("链接数据库成功")

}
