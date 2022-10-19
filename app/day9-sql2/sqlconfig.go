// 配置数据库连接

package app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// format an connection string
//
//root:123456@tcp(localhost:3306)/db
func SetConnString() (err error) {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", USERNAME, PASSWORD, ADDR, PORT, DATABASE)
	// TODO: 校验string是否正确
	if conn == "" {
		err = errors.New("the connection string is empty")
	}
	return
}

// 连接数据库
func SetConnection(conn string) error {

	// open connection to mysql server
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Printf("error opening connection to mysql server: %s", err)
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
	return err
}

// get db connection
func GetConnection() *sql.DB {

	return db

}

func CloseConnection() error {

	// TODO: close db connection
	return db.Close()
}
