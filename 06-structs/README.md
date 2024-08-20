## Structs

As structs em Go são tipos de dados compostos que permitem agrupar diferentes tipos de dados sob uma única estrutura. Elas são fundamentais para a organização e modelagem de dados em programas Go. Vamos explorar as structs em profundidade:

**Declaração de Structs**

A declaração de uma struct segue o seguinte formato:

```go
type NomeDaStruct struct {
    campo1 tipo1
    campo2 tipo2
    // ...
}

```

Por exemplo:

```go
type Pessoa struct {
    Nome    string
    Idade   int
    Endereco Endereco
}

```

**Acesso aos Campos**

Os campos de uma struct podem ser acessados usando a notação de ponto (.):

```go
var p Pessoa
p.Nome = "João"
p.Idade = 30

```

**Campos Embutidos**

É possível embutir uma struct em outra, o que permite acesso direto aos campos embutidos:

```go
type Endereco struct {
    Rua  string
    CEP  string
    Cidade string
}

type Pessoa struct {
    Nome    string
    Idade   int
    Endereco // Campos embutidos
}
```

Nesse caso, os campos da struct Endereco são acessados diretamente a partir de uma variável do tipo Pessoa.

**Métodos em Structs**

Em Go, é possível definir métodos associados a structs. Isso permite que você adicione comportamento específico a uma struct. A definição de método segue este padrão:

```go
func (receiver TipoReceptor) NomeDoMetodo(parametros) tipoRetorno {
    // corpo do método
}

```

Por exemplo:

```go
func (p *Pessoa) Envelhecer() {
    p.Idade++
}

```

**Métodos com Ponteiros vs. Valores**

Você pode definir métodos associados a structs usando ponteiros ou valores. Usar um ponteiro como receptor permite que você modifique diretamente o valor da struct dentro do método.

**Comparação de Structs**

Em Go, você pode comparar structs usando o operador de igualdade (==). As structs são consideradas iguais se todos os seus campos forem iguais. No entanto, structs não podem ser comparadas se contiverem campos que não podem ser comparados, como mapas, fatias ou funções.

**Tags de Estrutura**

As tags de estrutura (struct tags) permitem adicionar metadados aos campos de uma struct. Elas são frequentemente usadas para serialização e desserialização de dados.

**Structs Anônimas**

Você pode definir structs anonimamente, sem atribuir um nome a elas. Isso é útil ao definir structs dentro de outros tipos, como slices, maps ou como campos embutidos.
