// 配置数据库连接

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 声明db connection
var db *sql.DB

// 声明数据库连接常量
const (
	USERNAME     string = "root"
	PASSWORD     string = "123456"
	ADDR         string = "127.0.0.1"
	PORT         int    = 3306
	DATABASE     string = "wood_db1"
	MAXLIFETIME  int    = 10
	MAXIDLECONNS int    = 5
	MAXOPENCONNS int    = 10
)

// 连接数据库
func main() {

	// format an connection string
	//root:123456@tcp(localhost:3306)/db
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", USERNAME, PASSWORD, ADDR, PORT, DATABASE)

	// open connection to mysql server
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Printf("error opening connection to mysql server: %s", err)
		return
	}

	// set connection properties
	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxIdleConns(MAXIDLECONNS)
	db.SetMaxOpenConns(MAXOPENCONNS)

	// test connection
	err = db.Ping()
	if err != nil {
		// TODO: error handling
		log.Print(err)
	}

}
