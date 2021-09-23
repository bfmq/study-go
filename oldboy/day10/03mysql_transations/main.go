package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id       int
	username string
	password string
}

var db *sql.DB

func transactionDemo()  {
	tx,err :=db.Begin()
	if err != nil {
		fmt.Printf("begin failed,err:%s\n", err)
		return
	}
	tx.Rollback()
	tx.Commit()
}


func initDB() (err error) {
	dsn := "root:nrXTTQguLpXJGIkrExklwecwiQGQsjqYM1OVM35YHf12aK6LE5EnlopeWrno@tcp(10.10.2.20:3306)/go_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil { // 只是验证dsn格式
		fmt.Printf("dsn:%s invalid,err:%s\n", dsn, err)
		return
	}

	err = db.Ping() // 尝试链接数据库
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%s\n", dsn, err)
		return
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%s\n", err)
		return
	}
	fmt.Println("链接数据库成功")

	db.Close()
}
