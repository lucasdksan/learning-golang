## Loop

**Loop for**

O loop for é usado para repetir um bloco de código enquanto uma condição específica for verdadeira. A sintaxe básica é:

```go
for inicialização; condição; pós-loop {
    // código a ser repetido
}
```

Por exemplo:

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

O for em Go é muito flexível. Você pode omitir a inicialização e/ou a condição e/ou o pós-loop, criando diferentes tipos de loops.

**Loop for Infinito**

Você pode criar um loop infinito omitindo todas as três partes do for:

```go
for {
    // código a ser repetido indefinidamente
}
```

**Loop for com Range**

O loop for com range é usado para iterar sobre slices, arrays, strings, maps e canais. A sintaxe é:

```go
for indice, valor := range colecao {
    // código a ser repetido para cada elemento da coleção
}
```

Por exemplo:

```go
numeros := []int{1, 2, 3, 4, 5}
for indice, valor := range numeros {
    fmt.Println("Índice:", indice, "Valor:", valor)
}
```

**Uso de break e continue**

break: Termina o loop atual imediatamente.
continue: Pula a execução das iterações restantes do loop atual e vai para a próxima iteração.

Por exemplo:

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Sai do loop quando i for igual a 5
    }
    if i%2 == 0 {
        continue  // Pula as iterações pares
    }
    fmt.Println(i)
}
```
