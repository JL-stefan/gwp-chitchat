package data

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/chitchat?charset=utf8")
	if err != nil{
		log.Fatal(err)
	}
	return
}