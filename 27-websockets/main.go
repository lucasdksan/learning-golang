package main

import (
	"log"
	"net/http"

	"github.com/coder/websocket"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})

	if err != nil {
		log.Fatal(err)
		return
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
			break
		}

		log.Print("Recebido: ", string(data))
		c.Write(r.Context(), websocket.MessageText, data)
	}
}

func main() {
	http.HandleFunc("/", wsHandler)

	http.ListenAndServe(":8080", nil)
}
