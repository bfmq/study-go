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

func queryOne(id int) {
	var u1 user
	sqlStr := `select id,username,password from linux_user where id = ?;`
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.username, &u1.password)
	fmt.Println(u1)
}

func queryMore(n int) {
	sqlStr := `select id,username,password from linux_user where id > ?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed,err:%s\n", sqlStr, err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.username, &u1.password)
		if err != nil {
			fmt.Printf("scan failed,err:%s\n", err)
			return
		}
		fmt.Println(u1)
	}
}

func initDB() (err error) {
	dsn := "root:@tcp(10.10.2.20:3306)/go_test"
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

	// queryOne(53)
	// queryMore(50)
	db.Begin()

	db.Close()
}
