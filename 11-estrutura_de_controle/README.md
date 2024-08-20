## Estruturas de Controle

**Estrutura if**

A estrutura if é usada para executar um bloco de código se uma condição for avaliada como verdadeira. A sintaxe básica é a seguinte:

```go
if condição {
    // código a ser executado se a condição for verdadeira
}
```

Por exemplo:

```go
idade := 18
if idade >= 18 {
    fmt.Println("Você é maior de idade.")
}
```

Você também pode ter uma instrução if com uma inicialização de variável, onde a variável declarada é válida apenas no escopo do if:

```go
if idade := 18; idade >= 18 {
    fmt.Println("Você é maior de idade.")
}
```

**Estrutura if e else**

Além do if, você também pode usar o else para executar um bloco de código se a condição não for verdadeira. A sintaxe básica é:

```go
if condição {
    // código a ser executado se a condição for verdadeira
} else {
    // código a ser executado se a condição não for verdadeira
}
```

Por exemplo:

```go
    idade := 16
    if idade >= 18 {
        fmt.Println("Você é maior de idade.")
    } else {
        fmt.Println("Você é menor de idade.")
    }
```

**Estrutura if e else if**

Você pode encadear múltiplas condições usando else if para testar várias condições em sequência. A sintaxe é:

```go
    if condição1 {
        // código a ser executado se a condição1 for verdadeira
    } else if condição2 {
        // código a ser executado se a condição2 for verdadeira
    } else {
        // código a ser executado se nenhuma das condições anteriores for verdadeira
    }
```
Por exemplo:

```go
    idade := 16
    if idade >= 18 {
        fmt.Println("Você é maior de idade.")
    } else if idade >= 16 {
        fmt.Println("Você está quase lá.")
    } else {
        fmt.Println("Você é menor de idade.")
    }
```
