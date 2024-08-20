## Declaração e Inicialização

Em Go, você pode declarar uma variável usando a sintaxe var nome tipo. Por exemplo:

```go
    var idade int
```

Isso declara uma variável chamada idade do tipo int, que é um tipo de dado numérico inteiro em Go. Você também pode inicializar a variável ao mesmo tempo:

```go
    var idade int = 25
```

Ou, se você omitir o tipo, o compilador Go inferirá automaticamente o tipo com base no valor fornecido:

```go
    var idade = 25
```

Inferência de Tipo, Go permite a inferência de tipo, onde o tipo da variável é inferido a partir do valor atribuído a ela. Por exemplo:

```go
    nome := "João"  // O tipo de 'nome' é inferido como string
```

Zero Values, se uma variável é declarada, mas não inicializada explicitamente, ela recebe o valor zero de acordo com o tipo. Por exemplo:

```go
    var altura float64  // altura é inicializada com 0.0
```

**Atribuição de Valor**

Você pode atribuir um valor a uma variável em Go usando o operador =. Por exemplo:

```go
    idade = 30
```

Isso atribui o valor 30 à variável idade. Em Go, todas as variáveis devem ser usadas após serem declaradas ou inicializadas, caso contrário, o compilador emitirá um erro.