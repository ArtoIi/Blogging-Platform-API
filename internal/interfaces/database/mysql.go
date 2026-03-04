package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection(user, password, host, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, host, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
