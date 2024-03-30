package main

import (
	"crud/server"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD -> Create, Read, Update, Delete
	router := mux.NewRouter()
	router.HandleFunc("/users", server.Create_User).Methods(http.MethodPost)

	fmt.Println("Server on!")
	http.ListenAndServe(":5000", router)
}
