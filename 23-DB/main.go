package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	string_conection := "root:ss123456789@tcp(localhost:3306)/xeretech?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", string_conection)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão aberta!")

	lines, err := db.Query("select * from users")

	if err != nil {
		log.Fatal(err)
	}

	defer lines.Close()

	fmt.Println(lines)
}
