package app

import "database/sql"

// 声明db connection

type User struct {
	id       int
	username string
	password string
}

var (
	u  User
	db *sql.DB
)

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
