package app

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func PrepareInsert() {

	// 预处理statement
	stmt, err := db.Prepare("INSERT INTO sys_user(username,password) values (?,?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmt.Close()

	// 操作statement
	_, err = stmt.Exec("admin1", "123456")
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec("admin2", "12345")
	if err != nil {
		log.Fatal(err)
		return
	}

}
func PrepareUpdate() {

}
func PrepareDelete() {

}
