## WebSocket

O WebSocket é um protocolo de comunicação bidirecional que permite a comunicação em tempo real entre um cliente (como um navegador web) e um servidor. Ao contrário do protocolo HTTP, que é baseado em requisições e respostas, o WebSocket mantém uma conexão aberta entre o cliente e o servidor, permitindo que dados sejam enviados e recebidos em ambas as direções sem a necessidade de iniciar uma nova conexão a cada comunicação.

**Características principais do WebSocket:**

* Conexão persistente: Depois que a conexão é estabelecida, ela permanece aberta, permitindo comunicação contínua.

* Baixa latência: Por não exigir a reabertura da conexão para cada mensagem, a latência é muito menor em comparação com o HTTP.

* Bidirecionalidade: Tanto o cliente quanto o servidor podem enviar dados a qualquer momento, sem a necessidade de esperar por uma requisição.

* Compatibilidade com navegadores modernos: A maioria dos navegadores modernos suporta o protocolo WebSocket, o que o torna uma escolha viável para aplicações web que exigem atualizações em tempo real.

**Casos de uso comuns:**

* Aplicações de chat: Permite a troca de mensagens em tempo real entre usuários.
* Atualizações em tempo real: Usado para dashboards que precisam exibir dados atualizados continuamente, como preços de ações ou placares de esportes.
* Jogos online: Facilita a comunicação entre jogadores e servidores, oferecendo uma experiência de jogo sem atrasos.
* Notificações: Pode ser usado para enviar notificações instantâneas para os usuários.

**Coder Websocket**

O pacote github.com/coder/websocket é uma implementação leve e eficiente do protocolo WebSocket para a linguagem Go (Golang). Ele oferece uma API simples para adicionar suporte a WebSockets em aplicações Go, permitindo a comunicação em tempo real entre o servidor Go e os clientes conectados.

**Características principais do pacote coder/websocket:**

* Leve e eficiente: Focado na simplicidade e performance, é ideal para aplicações que necessitam de alta eficiência.

* API intuitiva: O pacote fornece uma API de fácil uso para manipular conexões WebSocket, incluindo suporte a operações de leitura e escrita de mensagens.

* Compatibilidade: Totalmente compatível com o padrão WebSocket, permitindo a integração com qualquer cliente que suporte o protocolo WebSocket.

* Suporte a múltiplas conexões: Permite gerenciar várias conexões WebSocket simultaneamente, o que é essencial para aplicações que precisam lidar com muitos clientes ao mesmo tempo.

**Exemplo básico de uso:**

Aqui está um exemplo básico de como usar o pacote github.com/coder/websocket para criar um servidor WebSocket simples em Go:

```go
package main

import (
    "log"
    "net/http"

    "github.com/coder/websocket"
)

func handler(w http.ResponseWriter, r *http.Request) {
    conn, err := websocket.Accept(w, r, nil)
    if err != nil {
        log.Println("Erro ao aceitar conexão:", err)
        return
    }
    defer conn.Close(websocket.StatusInternalError, "erro interno do servidor")

    for {
        _, msg, err := conn.Read(r.Context())
        if err != nil {
            log.Println("Erro ao ler mensagem:", err)
            return
        }
        log.Printf("Mensagem recebida: %s", msg)

        // Enviando uma resposta de volta ao cliente
        err = conn.Write(r.Context(), websocket.MessageText, msg)
        if err != nil {
            log.Println("Erro ao enviar mensagem:", err)
            return
        }
    }
}

func main() {
    http.HandleFunc("/ws", handler)
    log.Println("Servidor WebSocket iniciado na porta 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Erro ao iniciar servidor:", err)
    }
}
```

**Explicação do código:**

* Aceitação de conexão: No handler, o servidor aceita uma nova conexão WebSocket usando websocket.Accept. Se a conexão for aceita com sucesso, um objeto conn é retornado para gerenciar a comunicação com o cliente.

* Leitura de mensagens: O servidor lê mensagens do cliente com conn.Read. Aqui, as mensagens são recebidas em um loop, permitindo comunicação contínua.

* Escrita de mensagens: Após receber uma mensagem do cliente, o servidor a ecoa de volta usando conn.Write.

* Encerramento de conexão: A conexão é encerrada adequadamente usando conn.Close, fornecendo um código de status e uma razão para o encerramento.

**Casos de uso:**

* Aplicações de chat: Para permitir comunicação em tempo real entre usuários.

* Servidores de jogos: Para manter a comunicação contínua entre o servidor e os jogadores.

* Sistemas de notificação: Para enviar notificações em tempo real para clientes conectados.

* Monitoramento em tempo real: Ideal para dashboards que precisam de atualizações instantâneas.