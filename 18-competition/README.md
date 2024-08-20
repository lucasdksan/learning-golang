## Concorrência

Concorrência em Go é uma característica poderosa que permite que programas realizem múltiplas tarefas simultaneamente. Uma das principais ferramentas para lidar com concorrência em Go são as goroutines. Vamos explorar em profundidade como as goroutines funcionam e como são usadas para criar programas concorrentes eficientes.

**O que são Goroutines?**

Goroutines são leves threads gerenciadas pelo runtime de Go. Elas permitem que você execute funções de forma concorrente, de modo que várias tarefas possam ser realizadas simultaneamente dentro de um único programa Go. Goroutines são muito eficientes em termos de uso de memória e podem ser criadas e gerenciadas facilmente.

**Criando Goroutines**

Para criar uma goroutine, basta colocar a palavra-chave go antes de uma chamada de função. Isso instrui o runtime de Go a executar essa função em uma nova goroutine, simultaneamente com o restante do programa.

```go
func minhaFuncao() {
    // código da função
}

func main() {
    // Criando uma goroutine
    go minhaFuncao()

    // Outro código no programa principal
}
```

**Wait Group**

Em Go, o sync.WaitGroup é uma estrutura muito útil para sincronizar e coordenar várias goroutines. Ela é especialmente útil quando você precisa esperar que um grupo de goroutines seja concluído antes de continuar a execução do programa. Vamos explorar em profundidade como usar o sync.WaitGroup:

**O que é um WaitGroup?**

Um sync.WaitGroup é uma estrutura de dados fornecida pela biblioteca padrão de Go, localizada no pacote sync, que fornece uma maneira de esperar que um conjunto de operações assíncronas (goroutines) seja concluído antes de prosseguir com a execução do programa principal.

**Funcionamento Básico**

Adicionando Goroutines: Antes de iniciar uma goroutine que você deseja esperar, você precisa adicionar essa goroutine ao WaitGroup usando o método Add().

Sinalizando a Conclusão: Dentro de cada goroutine, você chama o método Done() do WaitGroup para indicar que a goroutine foi concluída.

Esperando a Conclusão: Depois de adicionar todas as goroutines ao WaitGroup e iniciar sua execução, você chama o método Wait() do WaitGroup no programa principal. Isso faz com que o programa principal aguarde até que todas as goroutines tenham sido concluídas, ou seja, até que o contador interno do WaitGroup retorne a zero.

**Exemplo de Uso**

Vamos dar uma olhada em um exemplo simples para entender melhor:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func minhaGoroutine(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Indica que a goroutine foi concluída ao sair da função
    fmt.Println("Goroutine", id, "iniciada")
    time.Sleep(time.Second) // Simula uma operação demorada
    fmt.Println("Goroutine", id, "concluída")
}

func main() {
    var wg sync.WaitGroup // Cria uma nova WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Adiciona uma goroutine ao WaitGroup
        go minhaGoroutine(i, &wg) // Inicia a goroutine
    }

    wg.Wait() // Aguarda a conclusão de todas as goroutines
    fmt.Println("Todas as goroutines concluídas")
}
```

**Canais**

Canais em Go são uma poderosa ferramenta de comunicação entre goroutines, permitindo a troca segura de dados e a sincronização de processos concorrentes. Vamos explorar em profundidade os canais em Go.

**O que são Canais?**

Canais em Go são estruturas de dados que fornecem uma maneira de uma goroutine enviar dados para outra goroutine de forma segura e eficiente. Eles facilitam a comunicação e a sincronização entre goroutines, tornando a concorrência mais segura e mais fácil de gerenciar.

**Declaração de Canais**

Você pode declarar um canal usando a função make() com a palavra-chave chan seguida do tipo de dados que será transmitido pelo canal:

```go
ch := make(chan int)  // Cria um canal para transmitir valores inteiros
```

**Operações em Canais**

Enviar Dados: Você pode enviar dados para um canal usando o operador <-.

```go
ch <- valor  // Envia o valor para o canal ch
```

Receber Dados: Você pode receber dados de um canal usando o operador <-. Esse operador atribui o valor recebido a uma variável.

```go
valor := <-ch  // Recebe um valor do canal ch e o atribui à variável valor
```

**Bloqueio de Canais**

Envio e Recebimento Bloqueantes: As operações de envio e recebimento em um canal são bloqueantes por padrão. Isso significa que uma goroutine que envia ou recebe dados de um canal será bloqueada até que a outra goroutine esteja pronta para receber ou enviar os dados, respectivamente.

**Bufferização de Canais**

Você pode criar canais com uma capacidade de buffer usando a função make() com um argumento de capacidade:

```go
ch := make(chan int, capacidade)  // Cria um canal com capacidade de buffer
```

Isso permite que o canal armazene vários valores em um buffer interno, permitindo que as goroutines enviem dados para o canal sem bloquear até que o buffer esteja cheio.

**Fechamento de Canais**

Você pode fechar um canal usando a função close(). Isso indica que não haverá mais valores enviados para o canal e que as goroutines que estão esperando para receber valores do canal devem terminar.

```go
close(ch)  // Fecha o canal ch
```

**Select**

O select é uma instrução de controle em Go que permite lidar com múltiplas operações de comunicação de forma concorrente. Ele é usado em conjunto com canais e permite que o programa escolha qual operação executar quando múltiplas estão prontas.

```go
select {
case msg1 := <-ch1:
    fmt.Println("Recebido", msg1)
case ch2 <- "mensagem":
    fmt.Println("Enviado mensagem")
}
```

**Worker Pools**

Os "worker pools" (piscinas de trabalhadores) são um padrão de concorrência comumente usado em programação concorrente para processamento paralelo de tarefas. Em Go, os worker pools são implementados usando goroutines e canais. Vamos explorar em profundidade como criar e usar worker pools em Go.

**O que são Worker Pools?**

Um worker pool é um conjunto de goroutines que estão prontas para receber e executar tarefas. Ele permite distribuir tarefas entre um grupo de goroutines para processamento paralelo, o que pode melhorar significativamente o desempenho e a eficiência do programa, especialmente para operações intensivas de CPU ou I/O.

**Componentes de um Worker Pool**

Um worker pool geralmente consiste em três componentes principais:

* Canal de Entrada (Input Channel): Este canal é usado para enviar tarefas para as goroutines do worker pool.

* Goroutines Trabalhadoras (Worker Goroutines): Estas são as goroutines responsáveis por executar as tarefas recebidas do canal de entrada.

* Canal de Saída (Output Channel): Opcionalmente, este canal pode ser usado para coletar os resultados das tarefas executadas pelas goroutines trabalhadoras.

**Implementação de um Worker Pool em Go**

Aqui está um exemplo simples de como implementar um worker pool em Go:

```go
package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Println("Worker", id, "processando o trabalho", job)
        results <- job * 2  // Simula um processamento e envia o resultado para o canal de saída
    }
}

func main() {
    numTrabalhadores := 3
    numTarefas := 5

    // Canais de entrada e saída
    jobs := make(chan int, numTarefas)
    results := make(chan int, numTarefas)

    // Iniciar goroutines trabalhadoras
    for i := 1; i <= numTrabalhadores; i++ {
        go worker(i, jobs, results)
    }

    // Enviar tarefas para o canal de entrada
    for j := 1; j <= numTarefas; j++ {
        jobs <- j
    }
    close(jobs)

    // Coletar resultados do canal de saída
    for k := 1; k <= numTarefas; k++ {
        result := <-results
        fmt.Println("Resultado:", result)
    }
}
```

Neste exemplo, criamos um worker pool com três goroutines trabalhadoras. As tarefas são enviadas para as goroutines através de um canal de entrada, e os resultados são coletados de um canal de saída.

***Vantagens dos Worker Pools**

* Paralelismo Controlado: Você pode controlar o número de goroutines trabalhadoras no pool para evitar a sobrecarga do sistema.

* Reutilização de Goroutines: As goroutines no pool podem ser reutilizadas para processar várias tarefas, evitando o custo de criar e destruir goroutines repetidamente.

* Distribuição Equilibrada de Tarefas: As tarefas são distribuídas entre as goroutines de forma equilibrada, garantindo uma utilização eficiente dos recursos.

**Generator**

Em Go, os geradores (generators) são uma técnica para criar iteradores de forma concisa e eficiente. Embora Go não tenha suporte nativo para geradores como em algumas outras linguagens, é possível implementar padrões semelhantes usando goroutines e canais. Vamos explorar como implementar geradores em Go.

**O que são Geradores?**

Um gerador é uma função que pode ser pausada e retomada, produzindo uma sequência de valores ao longo do tempo. Eles são úteis para gerar grandes conjuntos de dados de forma preguiçosa, economizando memória e CPU.

**Implementação de Geradores em Go**

Em Go, os geradores podem ser implementados usando goroutines e canais. Aqui está um exemplo simples de como criar um gerador em Go:

```go
package main

import "fmt"

func generateNumbers(max int) <-chan int {
    ch := make(chan int)

    go func() {
        defer close(ch)
        for i := 0; i < max; i++ {
            ch <- i
        }
    }()

    return ch
}

func main() {
    // Criar um gerador de números
    numbers := generateNumbers(5)

    // Iterar sobre os números gerados
    for num := range numbers {
        fmt.Println(num)
    }
}
```

Neste exemplo, a função generateNumbers retorna um canal de inteiros que produz uma sequência de números de 0 até max - 1. A goroutine interna gera os números e os envia para o canal, que é então iterado na função main.

**Características dos Geradores em Go**

* Preguiçoso (Lazy): Os geradores em Go são preguiçosos, o que significa que os valores são produzidos sob demanda à medida que são necessários, em vez de serem gerados todos de uma vez.

* Eficiente: Usar canais e goroutines para implementar geradores em Go é eficiente em termos de consumo de memória e CPU, pois permite que os valores sejam gerados de forma assíncrona e de maneira incremental.

**Uso de Geradores**

Os geradores são úteis sempre que você precisar gerar grandes conjuntos de dados de forma preguiçosa ou quando precisar processar dados de entrada em lote.

**Multiplexador**

Um multiplexador, em Go, geralmente se refere a um padrão de concorrência em que várias goroutines de entrada são combinadas em uma única goroutine de saída. Isso é feito utilizando canais para enviar dados de várias goroutines para um único ponto de processamento. Vamos explorar em profundidade como implementar um multiplexador em Go.

**O que é um Multiplexador?**

Um multiplexador é um dispositivo ou padrão de programação que combina múltiplas entradas em uma única saída. Em programação concorrente, um multiplexador é usado para processar dados de várias goroutines de entrada em uma única goroutine de saída.

**Implementação de um Multiplexador em Go**

Aqui está um exemplo simples de como implementar um multiplexador em Go:

```go
package main

import (
    "fmt"
)

func input1(ch chan<- string) {
    ch <- "Dados da fonte 1"
}

func input2(ch chan<- string) {
    ch <- "Dados da fonte 2"
}

func multiplex(input1 <-chan string, input2 <-chan string, output chan<- string) {
    for {
        select {
        case data := <-input1:
            output <- data
        case data := <-input2:
            output <- data
        }
    }
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    output := make(chan string)

    // Iniciar goroutines de entrada
    go input1(ch1)
    go input2(ch2)

    // Iniciar goroutine multiplexadora
    go multiplex(ch1, ch2, output)

    // Receber dados da saída
    data := <-output
    fmt.Println("Dados recebidos:", data)
}
```

Neste exemplo, duas goroutines de entrada (input1 e input2) enviam dados para um canal de saída compartilhado. A goroutine multiplex recebe dados de ambas as goroutines de entrada usando select e os envia para o canal de saída.

**Características de um Multiplexador em Go**

* Combinar Múltiplas Fontes: Um multiplexador em Go permite combinar dados de múltiplas fontes em uma única saída.

* Processamento Assíncrono: As goroutines de entrada e o multiplexador operam de forma assíncrona, o que significa que não há bloqueio enquanto aguardam dados.

* Sincronização de Dados: O uso de canais em Go para enviar e receber dados garante que as operações de leitura e escrita sejam seguras e sincronizadas.

**Uso de Multiplexadores**

Os multiplexadores são úteis sempre que você precisa combinar dados de várias fontes em uma única saída ou quando precisa processar dados de forma assíncrona e concorrente.
