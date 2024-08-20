## Contexts

Em Go, o conceito de contexto (context) é amplamente utilizado para gerenciar a execução de operações e controlar a propagação de valores, sinais de cancelamento e deadlines através de uma aplicação. A biblioteca padrão do Go oferece o pacote context, que facilita a criação de fluxos de trabalho que precisam respeitar prazos, realizar operações assíncronas ou interromper a execução sob determinadas condições. 

**Context de forma resumida**

Contextos basicamente utilizam channels por debaixo dos panos, nos quais, os processos que os utilizam precisam ficar ouvindo os sinais que recebem deste contexto em questão.

Abaixo estão alguns dos principais contextos em Go:

1. Contextos Básicos

* context.Background(): É o contexto vazio inicial, normalmente utilizado no início de um fluxo de operações. Ele não carrega nenhuma informação e nunca é cancelado ou tem um deadline. É comum vê-lo usado como contexto raiz.

* context.TODO(): Similar ao Background(), mas é usado quando você ainda não sabe qual contexto utilizar. Indica que algo precisa ser definido futuramente.

2. Propagação de Cancelamento

* context.WithCancel(parent): Cria um novo contexto derivado de um contexto pai, que pode ser cancelado explicitamente. Ao cancelar este contexto, todos os contextos derivados dele também serão cancelados, permitindo interromper operações em cascata.

*Observação: Podemos mandar sinais para nosso código terminar a execução imediatamente, matando assim o processo sendo executado, seja síncrono ou assíncrono.*

3. Deadline e Timeout

* context.WithDeadline(parent, deadline): Cria um novo contexto que automaticamente será cancelado após atingir um tempo específico (deadline). É útil para operações que devem ser concluídas dentro de um limite de tempo.

* context.WithTimeout(parent, timeout): Similar ao WithDeadline, mas define o tempo a partir do momento da criação do contexto. Útil para definir timeouts específicos em operações.

4. Armazenamento de Valores

* context.WithValue(parent, key, value): Permite associar valores específicos ao contexto, que podem ser recuperados por funções subsequentes que recebem o mesmo contexto. No entanto, deve ser usado com cautela, pois pode tornar o código mais difícil de entender e manter.

5. Utilização no Go

* Gerenciamento de Goroutines: Contextos são frequentemente usados para gerenciar goroutines, permitindo cancelar múltiplas goroutines que estão sendo executadas paralelamente.

* HTTP Requests: No contexto de servidores web, o pacote net/http do Go é integrado com o context para garantir que quando uma requisição é cancelada ou times out, todos os processos relacionados àquela requisição também sejam cancelados.

* Banco de Dados e I/O: Em operações de I/O ou chamadas a banco de dados, o contexto permite definir timeouts, o que é crucial para evitar que a aplicação fique bloqueada em operações demoradas.
