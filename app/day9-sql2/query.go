package app

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func QueryRow() {
	err := db.QueryRow("select id, username, password from 'sys_user' where id > ?", 2).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func QueryMultipleRow() {

	// 获取多行结果
	rows, err := db.Query("select id, username, password from 'sys_user' where id > ?", 2)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 关闭rows, 释放连接
	defer rows.Close()

	// 循环打印每一行的数据
	for rows.Next() {
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("userid:%v\n, username:%v\n, userpassword:%v\n", u.id, u.username, u.password)
	}

}

func PrepareQuery() {

	// 预处理statement语句，分离变量防止SQL注入
	stmt, err := db.Prepare("select id, username, password from 'sys_user' where id > ?")
	if err != nil {
		log.Fatal(err)
	}

	// 关闭stmt, 释放连接
	defer stmt.Close()

	rows, err := stmt.Query(2)
	if err != nil {
		log.Fatal(err)
	}

	// 关闭rows，释放连接
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("userid:%v\n, username:%v\n, userpassword:%v\n", u.id, u.username, u.password)
	}

}
