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
	router.HandleFunc("/users", server.Search_All_Users).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.Search_User).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.Update_User).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", server.Delete_User).Methods(http.MethodDelete)

	fmt.Println("Server on!")
	http.ListenAndServe(":5000", router)
}
