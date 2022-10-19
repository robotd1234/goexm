package app

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Transaction() {

	// db begin
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return
	}

	// execute transaction

	tx.Prepare("INSERT INTO sys_user(username,password) values (?,?)")
	_, err = tx.Exec("admin, user@123")
	if err != nil {
		log.Fatal(err)
		// tx rollback if fail
		tx.Rollback()
		return
	}
	// tx commit
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return
	}

}
