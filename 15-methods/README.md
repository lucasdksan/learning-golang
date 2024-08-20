## Métodos

Métodos em Go são funções associadas a um tipo específico, permitindo que você adicione comportamento específico aos valores desse tipo. Vou explicar em detalhes como os métodos funcionam em Go.

**Declaração de Métodos**

A declaração de um método em Go segue o seguinte formato:

```go
func (receptor TipoReceptor) NomeDoMetodo(parametros) tipoRetorno {
    // corpo do método
}
```

receptor: É o nome dado ao valor do tipo para o qual o método está sendo definido. O receptor é colocado entre parênteses antes do nome do método.

TipoReceptor: É o tipo de dados para o qual o método está sendo definido.

NomeDoMetodo: É o nome do método.

parametros: São os parâmetros que o método recebe.

tipoRetorno: É o tipo de retorno do método.

**Exemplo de Método**

Vamos criar um método Area() para o tipo Retangulo que calcula a área do retângulo:

```go
type Retangulo struct {
    Largura  float64
    Altura float64
}

func (r Retangulo) Area() float64 {
    return r.Largura * r.Altura
}
```

**Chamada de Método**

Os métodos são chamados usando a notação de ponto (.) no valor do tipo associado:

```go
retangulo := Retangulo{Largura: 10, Altura: 5}
area := retangulo.Area() // Chamada do método
fmt.Println("Área do retângulo:", area)
```

**Receptores Ponteiros vs. Valores**

Em Go, você pode definir métodos para tipos de receptor ponteiro ou tipos de valor. Quando um método é definido com um receptor ponteiro, as alterações feitas dentro do método afetam diretamente o valor original, enquanto com um receptor valor, uma cópia do valor é passada para o método.

**Métodos Anexados a Tipos Embutidos**

É possível anexar métodos a tipos embutidos (como structs embutidas ou tipos básicos como int, float, etc.). Isso pode ser útil para adicionar métodos a tipos que você não pode modificar diretamente.

**Métodos como Funções**

Em Go, os métodos são basicamente funções, mas com um argumento especial (o receptor) que é passado implicitamente. Isso torna os métodos uma forma elegante de adicionar comportamento aos tipos em Go.
