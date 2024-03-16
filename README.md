# learning-golang

*Meus primeiros passos em GOlang*

## História

Go, também conhecida como Golang, é uma linguagem de programação de código aberto criada pelos engenheiros do Google Robert Griesemer, Rob Pike e Ken Thompson. O desenvolvimento começou em 2007, com o objetivo de criar uma linguagem que fosse eficiente, concorrente, segura e de fácil manutenção.

A motivação para criar o Go surgiu da necessidade do Google de uma linguagem que pudesse lidar com sistemas de larga escala, especialmente em servidores web e infraestrutura de software. O Go foi lançado publicamente em 2009 e desde então tem crescido em popularidade, sendo amplamente utilizado em diversos domínios, incluindo desenvolvimento de sistemas distribuídos, computação em nuvem, microserviços e desenvolvimento de aplicações web.

O Go foi projetado com foco na simplicidade e na produtividade dos desenvolvedores. Ele possui uma sintaxe limpa e concisa, suporta concorrência nativa através de goroutines e canais, o que facilita a construção de sistemas paralelos e distribuídos. Além disso, o sistema de gerenciamento de pacotes integrado (o go get), a coleta de lixo eficiente e a compilação rápida são características que contribuem para a popularidade da linguagem.

Atualmente, o Go é mantido pela comunidade open source e recebe contribuições de desenvolvedores de todo o mundo. Sua evolução continua, com melhorias constantes na linguagem, bibliotecas e ferramentas relacionadas.

**Extremamente Eficiente.**

* Compilado;
* Otimizada para usar mais de um núcleo do processador;
* Feita para lidar muito bem com CONCORRÊNCIA;

## Importação e Exportação de Funções e Pacotes em Go

Em Go, a importação e exportação de funções e pacotes são fundamentais para construir aplicativos modulares e reutilizáveis. Aqui estão os principais pontos sobre como isso funciona:

**Exportação de Funções e Identificadores**

Em Go, a exportação de funções e identificadores é feita através da primeira letra maiúscula. Se um identificador (nome de função, tipo, variável, etc.) começa com uma letra maiúscula, ele é exportado e pode ser acessado por outros pacotes. Por exemplo:

```go
    package math

    func Add(a, b int) int {
        return a + b
    }
```

A função Add é exportada porque começa com uma letra maiúscula.

**Importação de Pacotes**

Para importar um pacote em Go, você usa a declaração import seguida pelo caminho do pacote. Por exemplo:

```go
    import "fmt"
    import "math"
```

**Importação de Pacotes da Internet**

Go torna fácil importar pacotes da internet. Normalmente, você usará um sistema de controle de versão como Git para hospedar seus pacotes. Por exemplo, se você tem um pacote chamado github.com/user/package, você pode importá-lo em seu código assim:

```go
    import "github.com/user/package"
```

Quando você compila seu código, o compilador Go automaticamente baixa e instala o pacote do repositório Git especificado.

Se você está desenvolvendo um pacote que deseja disponibilizar para outros desenvolvedores, você pode usar o sistema de módulos Go. Com o sistema de módulos, você pode especificar dependências em um arquivo go.mod e importar pacotes usando nomes de módulos.

Por exemplo, se você estiver desenvolvendo um módulo chamado example.com/mymodule, e quiser importar o pacote github.com/user/package, você pode adicioná-lo ao seu go.mod e importá-lo em seu código como qualquer outro pacote:

```go
    import "github.com/user/package"    
```

E no seu go.mod:

```go

    module example.com/mymodule

    require github.com/user/package v1.2.3
```

**Importação e Uso de Tipos e Funções**

Dentro do mesmo pacote, você pode simplesmente importar outros arquivos do mesmo pacote e usar os tipos e funções definidos nesses arquivos. Por exemplo, suponha que você tenha dois arquivos, file1.go e file2.go, ambos pertencentes ao mesmo pacote mypackage:

```go
// file1.go
package mypackage

type MyType struct {
    // campos e métodos aqui
}
```

```go
// file2.go
package mypackage

func MyFunction() {
    // código da função aqui
}
```

Você pode então importar e usar MyType e MyFunction em qualquer outro arquivo dentro do pacote mypackage.

```go
// outro_arquivo.go
package mypackage

func outraFuncao() {
    obj := MyType{}    // Usando MyType
    MyFunction()       // Chamando MyFunction
}
```

**Acesso a Variáveis Globais**

Dentro do mesmo pacote, você também pode acessar variáveis globais definidas em outros arquivos do mesmo pacote. No entanto, é importante notar que as variáveis globais devem ser exportadas (iniciar com letra maiúscula) para serem acessíveis de outros arquivos.

**Uso de Constantes e Variáveis Não Exportadas**

Mesmo que uma constante ou variável não seja exportada (ou seja, comece com letra minúscula), ela ainda pode ser usada em outros arquivos do mesmo pacote. Isso é útil quando você deseja manter constantes ou variáveis compartilhadas dentro do pacote, mas não quer que elas sejam acessíveis fora do pacote.

**Organização de Código**

A organização de código em arquivos separados dentro do mesmo pacote é uma prática comum em Go, especialmente para pacotes maiores. Isso ajuda a manter o código limpo e organizado, facilitando a manutenção e a colaboração entre os membros da equipe.

Em resumo, dentro do mesmo pacote em Go, você pode interagir com arquivos de várias maneiras, incluindo importação e uso de tipos, funções, variáveis globais, constantes e variáveis não exportadas. Isso oferece flexibilidade e organização ao desenvolver aplicativos em Go.

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

**Tipos de Variáveis em Go**

Go possui uma série de tipos de variáveis embutidos, incluindo:

int, int8, int16, int32, int64: tipos inteiros com diferentes tamanhos.
uint, uint8, uint16, uint32, uint64: tipos inteiros sem sinal com diferentes tamanhos.
float32, float64: tipos de ponto flutuante para números decimais.
bool: tipo booleano para representar verdadeiro ou falso.
string: tipo para representar cadeias de caracteres.
byte (alias para uint8) e rune (alias para int32): tipos especiais para representar bytes e caracteres Unicode, respectivamente.
Além desses tipos básicos, Go também permite a definição de tipos de variáveis personalizados usando structs, arrays, slices, maps e interfaces.

**Escopo de Variáveis**

Em Go, o escopo de uma variável é determinado pelo bloco em que é declarada. Por exemplo:

```go
func exemplo() {
    var x int = 10
    // x está disponível aqui
}
// x não está disponível aqui
```

Variáveis declaradas em um bloco interno têm precedência sobre variáveis com o mesmo nome em blocos externos. No entanto, se você declarar uma variável com o mesmo nome dentro de um bloco interno, ela só será válida dentro desse bloco e não afetará a variável externa com o mesmo nome.