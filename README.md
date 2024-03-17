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

## Tipos de Variáveis em Go

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

A tipagem em Go é uma característica fundamental da linguagem que oferece segurança e flexibilidade no desenvolvimento de software. Vou abordar diversos aspectos da tipagem em Go, desde sua natureza estática até suas implicações na segurança e na eficiência do código.

**Tipagem Estática**

Go é uma linguagem tipada estaticamente, o que significa que o tipo de cada variável é conhecido em tempo de compilação. Isso contrasta com linguagens dinâmicas, onde o tipo é determinado em tempo de execução. A tipagem estática traz benefícios significativos, como detecção precoce de erros de tipo e otimizações de desempenho.

**Inferência de Tipo**

Apesar de ser tipada estaticamente, Go suporta inferência de tipo, o que permite declarar variáveis sem especificar explicitamente seus tipos. O compilador deduz o tipo com base no valor atribuído. Por exemplo:

```go
    nome := "João"  // O tipo de 'nome' é inferido como string
```

**Tipos Básicos e Compostos**

Go oferece uma variedade de tipos básicos, como inteiros, ponto flutuante, booleanos e strings, bem como tipos compostos, como arrays, slices, maps e structs. Esses tipos fornecem a base para a construção de estruturas de dados complexas e a implementação de algoritmos eficientes.

**Segurança de Tipos**

A tipagem estática em Go contribui para a segurança do tipo, garantindo que operações inválidas, como adicionar um inteiro a uma string, sejam detectadas em tempo de compilação. Isso ajuda a evitar erros comuns e torna o código mais confiável e previsível.

**Polimorfismo**

Go suporta polimorfismo por meio de interfaces. Interfaces em Go permitem que os tipos sejam tratados de maneira genérica, possibilitando a criação de código flexível e reutilizável. Através do conceito de interfaces, Go promove a composição sobre a herança, o que simplifica o design e evita problemas comuns associados à herança múltipla.

**Eficiência de Tipo**

A tipagem estática em Go também contribui para a eficiência do código, permitindo ao compilador otimizar operações de acesso a memória e execução de instruções. O conhecimento antecipado dos tipos permite ao compilador gerar código altamente otimizado, resultando em desempenho superior em comparação com linguagens dinâmicas.

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

**Funções Anônimas**

Em Go, é possível criar funções sem nome, conhecidas como funções anônimas ou closures. Essas funções são frequentemente utilizadas em situações onde uma função é passada como argumento para outra função, como em goroutines ou no uso de funções de ordem superior. Por exemplo:

```go
    func main() {
        minhaFuncao := func() {
            fmt.Println("Olá, mundo!")
        }
        minhaFuncao()  // Chamada da função anônima
    }
```

**Closures**

As closures em Go são funções que capturam variáveis do contexto ao seu redor. Elas são particularmente úteis para criar funções que podem ser configuradas com diferentes contextos, como no padrão de fábrica de funções. Por exemplo:

```go
    func gerarIncrementador(incremento int) func() int {
        contador := 0
        return func() int {
            contador += incremento
            return contador
        }
    }
```

Neste exemplo, a função gerarIncrementador retorna uma closure que captura a variável incremento e retorna uma função que incrementa um contador interno.

**Funções Variádicas**

Em Go, é possível definir funções variádicas, que podem receber um número variável de argumentos. Para criar uma função variádica, você usa ... seguido do tipo do argumento variádico. Por exemplo:

```go
    func somarNumeros(nums ...int) int {
        total := 0
        for _, num := range nums {
            total += num
        }
        return total
    }
```

**Funções Recursivas**

Go suporta a recursão, permitindo que uma função chame a si mesma. No entanto, é importante considerar a eficiência e a base de casos da recursão para evitar estouro de pilha. Por exemplo:

```go
    func fatorial(n int) int {
        if n == 0 {
            return 1
        }
        return n * fatorial(n-1)
    }
```

**Concorrência e Goroutines**

Em Go, as goroutines são leves e podem ser iniciadas facilmente com a palavra-chave go. Elas são frequentemente utilizadas para realizar operações concorrentes de forma eficiente. Por exemplo:

```go
    func main() {
        go minhaFuncaoConcorrente()
        // outras operações
    }

    func minhaFuncaoConcorrente() {
        // código executado em paralelo
    }
```

## Operadores

**Operadores Aritméticos**

Os operadores aritméticos em Go são usados para realizar operações matemáticas básicas. Os principais operadores aritméticos são:

* + : Adição
* - : Subtração
* * : Multiplicação
* / : Divisão
* % : Módulo (resto da divisão)

Por exemplo:

```go
a := 10
b := 5
soma := a + b      // 15
subtracao := a - b // 5
multiplicacao := a * b // 50
divisao := a / b   // 2
modulo := a % b    // 0
```

**Operadores de Comparação**

Os operadores de comparação são usados para comparar valores. Eles retornam um valor booleano (true ou false). Os principais operadores de comparação em Go são:

* == : Igual a
* != : Diferente de
* < : Menor que
* -> : Maior que
* <= : Menor ou igual a
* ->= : Maior ou igual a

Por exemplo:

```go
a := 10
b := 5
igual := a == b       // false
diferente := a != b   // true
menor := a < b        // false
maior := a > b        // true
menorOuIgual := a <= b  // false
maiorOuIgual := a >= b  // true
```

**Operadores Lógicos**

Os operadores lógicos em Go são usados para combinar expressões booleanas. Os principais operadores lógicos são:

&& : E lógico (AND)
|| : Ou lógico (OR)
! : Negação lógica (NOT)

Por exemplo:

```go
a := true
b := false
eLogico := a && b   // false
ouLogico := a || b  // true
negacao := !a       // false
```

**Operadores Bit a Bit**

Go também suporta operadores bit a bit, que operam em cada bit individual de um número. Os principais operadores bit a bit são:

* & : E bit a bit (AND)
* | : Ou bit a bit (OR)
* ^ : Ou exclusivo bit a bit (XOR)
* << : Deslocamento para a esquerda
* ->> : Deslocamento para a direita

Por exemplo:

```go
a := 5   // 101 em binário
b := 3   // 011 em binário

eBitABit := a & b    // 1 (001)
ouBitABit := a | b   // 7 (111)
ouExclusivo := a ^ b // 6 (110)
deslocamentoEsquerda := a << 1  // 10 (1010)
deslocamentoDireita := a >> 1   // 2 (10)
```

**Operadores de Atribuição**

Os operadores de atribuição em Go são usados para atribuir valores a variáveis. O operador de atribuição básico é =. Além disso, há operadores de atribuição combinados que realizam uma operação aritmética antes de atribuir o valor. Por exemplo:

* += : Adição e atribuição
* -= : Subtração e atribuição
* *= : Multiplicação e atribuição
* /= : Divisão e atribuição
* %= : Módulo e atribuição
* <<= : Deslocamento para a esquerda e atribuição
* ->>= : Deslocamento para a direita e atribuição
* &= : E bit a bit e atribuição
* |= : Ou bit a bit e atribuição
* ^= : Ou exclusivo bit a bit e atribuição

Por exemplo:

```go
a := 10
a += 5   // a agora é 15
```

**Operador Ternário (Não Suportado em Go)**

É importante observar que Go não tem um operador ternário como algumas outras linguagens (por exemplo, ? : em C/C++). Em vez disso, geralmente se usa a estrutura if ou uma expressão condicional para alcançar o mesmo resultado.