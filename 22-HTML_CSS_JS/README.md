## HTML, CSS, JavaScript

Para renderizar arquivos HTML em um servidor HTTP em Go, podemos usar o pacote html/template da biblioteca padrão. Este pacote nos permite criar modelos HTML dinâmicos e preencher esses modelos com dados antes de enviá-los como resposta HTTP. Vou mostrar um exemplo de como fazer isso:

Suponha que temos um arquivo HTML chamado index.html na mesma pasta que nosso arquivo Go:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Exemplo</title>
</head>
<body>
    <h1>Olá, {{.Name}}!</h1>
    <p>Idade: {{.Age}}</p>
</body>
</html>
```

Aqui está um exemplo de código Go que renderiza este arquivo HTML:

```go
package main

import (
    "html/template"
    "net/http"
)

type Person struct {
    Name string
    Age  int
}

func handler(w http.ResponseWriter, r *http.Request) {
    // Definir os dados a serem inseridos no template
    data := Person{Name: "João", Age: 30}

    // Parse o arquivo HTML
    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Renderize o template com os dados
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

Neste exemplo, definimos uma estrutura Person que contém o nome e a idade de uma pessoa. No manipulador HTTP, criamos uma instância dessa estrutura e a passamos para o template HTML. Usamos template.ParseFiles para carregar o arquivo HTML e tmpl.Execute para renderizá-lo com os dados fornecidos e enviá-lo como resposta HTTP.

Lembre-se de que o arquivo HTML deve estar acessível para o servidor Go e que você deve manipular quaisquer erros que ocorram durante o processamento do template.
