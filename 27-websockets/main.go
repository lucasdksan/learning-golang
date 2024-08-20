package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/coder/websocket"
)

var (
	clients map[*websocket.Conn]bool = make(map[*websocket.Conn]bool)
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// Cliente conectou com o websocket

	clients[c] = true

	for client := range clients {
		client.Write(r.Context(), websocket.MessageText, []byte("Um novo cliente está conectado!"))
	}

	defer func() {
		err := c.Close(websocket.StatusNormalClosure, "Conexão encerrada normalmente")
		if err != nil {
			log.Println("Erro ao fechar a conexão:", err)
		}
	}()

	for {
		_, data, err := c.Read(r.Context())

		if err != nil {
			log.Println("Erro ao ler dados:", err)
			delete(clients, c)
			break
		}

		log.Print("Recebido: ", string(data))
		for client := range clients {
			client.Write(r.Context(), websocket.MessageText, data)
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	http.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(strconv.Itoa(len(clients))))
	})

	http.ListenAndServe(":8080", nil)
}
