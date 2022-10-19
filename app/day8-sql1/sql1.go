package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/wood_db1")
	if err != nil {
		return err
	}
	// 链接数据库
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {

	fmt.Printf("before connecting %v", db)
	err := initDB()
	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		return
	}
	fmt.Println("Success")
	fmt.Printf("after connecting %v", db)
}
