package database

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)
func connection() *sql.DB {
	db, err := sql.Open("mysql", "root:singgih1@tcp(localhost:3306)/belajar_golang_db")  
	if err != nil {
		panic(err)
	}

	//pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
func GetConnection() *sql.DB {
	return connection()
}
