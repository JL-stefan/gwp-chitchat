package data

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	var err error
	DB, err = sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/chitchat?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return

}
