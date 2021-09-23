package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func Init(dns string) (err error) {
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		fmt.Printf("coon mysql failed,err:%s\n", err)
		return
	}

	// 查看是否连成功
	err = DB.Ping()
	if err != nil {
		fmt.Printf("test mysql failed,err:%s\n", err)
		return
	}
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(16)
	return nil
}
