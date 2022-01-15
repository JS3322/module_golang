package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 트랜잭션 시작
	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO test1 VALUES (?, ?)", 15, "Jack")
	if err != nil {
		log.Panic(err)
	}

	_, err = tx.Exec("INSERT INTO test2 VALUES (?, ?)", 15, "Data")
	if err != nil {
		log.Panic(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}
