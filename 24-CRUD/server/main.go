package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Create_User(w http.ResponseWriter, r *http.Request) {
	body_reques, err := ioutil.ReadAll(r.Body)
	var user user

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	if err = json.Unmarshal(body_reques, &user); err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	db, erro := database.Connect()

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
	}

	statement, erro := db.Prepare("insert into users (name, email) values (?, ?)")

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
	}

	defer statement.Close()

	insertion, err := statement.Exec(user.Name, user.Email)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
	}

	insert_id, err := insertion.LastInsertId()

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", insert_id)))
}
