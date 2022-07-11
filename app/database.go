package app

import (
	"database/sql"
	"fmt"
	"my-gram/helper"
	"time"
)

func NewDB() *sql.DB  {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/mygramdb")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	fmt.Println("Successfully Connect to Database")
	return db
}