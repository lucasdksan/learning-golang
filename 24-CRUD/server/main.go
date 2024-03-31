package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("insert into users (name, email) values (?, ?)")

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	defer statement.Close()

	insertion, err := statement.Exec(user.Name, user.Email)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	insert_id, err := insertion.LastInsertId()

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", insert_id)))
}

func Search_All_Users(w http.ResponseWriter, r *http.Request) {
	var users []user

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	defer db.Close()

	lines, err := db.Query("select * from users")

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	defer lines.Close()

	for lines.Next() {
		var user user

		if err := lines.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Falha ao ler o corpo da requisição"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}
}

func Search_User(w http.ResponseWriter, r *http.Request) {
	var user user
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	defer db.Close()

	line, err := db.Query("select * from users where id = ?", ID)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Falha ao ler o corpo da requisição"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}
}

func Update_User(w http.ResponseWriter, r *http.Request) {
	var user user
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	body_reques, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	if err = json.Unmarshal(body_reques, &user); err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func Delete_User(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("delete from users where id = ?")

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
