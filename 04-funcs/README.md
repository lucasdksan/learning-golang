## Funções

As funções em Go são um componente essencial da linguagem e são usadas para encapsular lógica, promover a reutilização de código e facilitar a organização do programa. Vou explorar diversos aspectos das funções em Go, desde sua sintaxe básica até conceitos mais avançados, como funções anônimas e closures.

**Sintaxe Básica**

A sintaxe básica de uma função em Go é relativamente simples:

```go
    func nomeDaFuncao(parametros) tipoDeRetorno {
        // corpo da função
        return valorDeRetorno
    }
```

* func: Palavra-chave usada para definir uma função.
* nomeDaFuncao: Nome da função.
* parametros: Lista de parâmetros que a função recebe.
* tipoDeRetorno: Tipo de dado retornado pela função.
* return: Palavra-chave utilizada para retornar valores da função.

**Parâmetros e Retornos Múltiplos**

Em Go, uma função pode receber zero ou mais parâmetros e pode retornar zero ou mais valores. Por exemplo:

```go
    func soma(a, b int) int {
        return a + b
    }
```

Esta função soma recebe dois parâmetros do tipo inteiro e retorna a soma desses dois valores como um inteiro.

Em GO, é permitido uma função ter mais de um retorno.

```go
    func MathematicalCalculations(a, b int) (int, int) {
        soma := a + b
        sub := a - b

        return soma, sub // metodo para ter dois ou mais retorno em uma função
    }

    result_2, result_3 := mathematicalcalculations.MathematicalCalculations(10, 5)
```

A função MathematicalCalculations recebe dois valores e aplica as operações de soma e subtração. Em seguida retorna os dois resultados;
