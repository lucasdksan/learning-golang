## Maps 

Os maps em Go são uma estrutura de dados muito útil e versátil, fornecendo uma maneira eficiente de armazenar pares de chave-valor. Vou detalhar os conceitos dos maps em Go para que você tenha uma compreensão abrangente de como eles funcionam.

**O que são Maps?**

Um map em Go é uma coleção não ordenada de elementos, onde cada elemento é um par de chave-valor único. As chaves são usadas para indexar e acessar os valores associados.

**Declaração de Maps**

A declaração de um map em Go segue o seguinte formato:

```go
var nomeDoMap map[TipoDaChave]TipoDoValor
```

Por exemplo:

```go
var idade map[string]int  // Um map onde as chaves são strings e os valores são inteiros
```

**Inicialização de Maps**

Você pode inicializar um map usando a função make() ou utilizando a sintaxe de inicialização de map:

```go
// Usando make()
idade := make(map[string]int)

// Usando a sintaxe de inicialização de map
idade := map[string]int{
    "João": 30,
    "Maria": 25,
}
```

**Acesso a Elementos**

Você pode acessar elementos de um map usando suas chaves:

```go
fmt.Println(idade["João"])  // Imprime o valor associado à chave "João"
```

Se a chave não existir no map, o valor retornado será o zero do tipo do valor.

**Adição e Atualização de Elementos**

Para adicionar ou atualizar um elemento em um map, basta atribuir um valor a uma chave:

```go
idade["Pedro"] = 35  // Adiciona um novo elemento ao map ou atualiza o valor se a chave já existir
```

**Remoção de Elementos**

Para remover um elemento de um map, use a função delete():

```go
delete(idade, "Maria")  // Remove o elemento com a chave "Maria" do map
```

**Verificação de Existência de Chave**

Você pode verificar se uma chave existe em um map usando a sintaxe de atribuição múltipla:

```go
valor, ok := idade["João"]
if ok {
    fmt.Println("O valor de João é:", valor)
} else {
    fmt.Println("João não está presente no map")
}
```

**Iteração sobre Maps**

Você pode iterar sobre os elementos de um map usando um loop for range:

```go
for chave, valor := range idade {
    fmt.Println(chave, ":", valor)
}
```

**Maps vs. Slices vs. Arrays**

Maps são úteis quando você precisa associar valores a chaves específicas e quando a ordem dos elementos não é importante.
Slices são usados quando você precisa de uma coleção ordenada de elementos e quando a indexação por posição é necessária.
Arrays são usados quando o tamanho dos dados é fixo e conhecido antecipadamente.
