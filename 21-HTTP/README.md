## HTTP

Trabalhar com HTTP em Go é simples e direto, graças à biblioteca padrão rica em recursos. Vamos explorar como criar um servidor HTTP básico usando o pacote net/http e, em seguida, veremos como podemos expandir suas capacidades para criar um servidor mais robusto.

**Criando um Servidor HTTP Básico**

Para criar um servidor HTTP básico em Go, precisamos seguir três etapas principais:

* Definir um manipulador para lidar com solicitações HTTP.
* Registrar o manipulador para um caminho específico.
* Iniciar o servidor HTTP.

Aqui está um exemplo simples de um servidor HTTP que responde com "Hello, World!" para todas as solicitações:

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

**Manipulador HTTP**

O manipulador handler neste exemplo é uma função que recebe dois parâmetros: http.ResponseWriter e *http.Request. O primeiro é usado para escrever a resposta HTTP de volta para o cliente, e o segundo contém informações sobre a solicitação HTTP recebida.

**Registrando o Manipulador**

Usamos http.HandleFunc para associar um manipulador a um padrão de rota. Neste caso, estamos associando o manipulador ao padrão de rota raiz ("/"), o que significa que ele responderá a todas as solicitações recebidas no servidor.

**Iniciando o Servidor**

Finalmente, usamos http.ListenAndServe para iniciar o servidor HTTP na porta especificada (8080 neste exemplo). Este método é bloqueante e só retorna se ocorrer um erro.

**Expansão do Servidor HTTP**

Para tornar o servidor HTTP mais robusto, podemos adicionar recursos adicionais, como roteamento avançado, middleware, tratamento de erros, suporte a TLS/SSL, autenticação e autorização, e muito mais. Isso pode ser feito usando bibliotecas de terceiros ou implementando esses recursos manualmente.
