package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	string_conection := "root:ss123456789@tcp(localhost:3306)/xeretech?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", string_conection)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
