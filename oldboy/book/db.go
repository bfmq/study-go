package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	addr := "root:admin@tcp(127.0.0.1:3306)/test"
	db, err = sqlx.Connect("mysql", addr)
	if err != nil {
		fmt.Printf(",err:%s\n", err)
		return
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)

	return
}

func queryAllBook() (bookList []*Book, err error) {
	sqlStr := "select id, title,price from book"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Printf("queryAllBook查询失败,err:%s\n", err)
		return
	}
	return
}

func insertBook(title string, price int64) (err error) {
	sqlStr := "insert book (title, price) values (?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Printf("insertBook插入失败,err:%s\n", err)
		return
	}
	return
}

func deleteBook(id int64) (err error) {
	sqlStr := "delete from book where id = ?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("deleteBook删除失败,err:%s\n", err)
		return
	}
	return
}
