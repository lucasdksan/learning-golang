## Switch

**Estrutura switch**

A estrutura switch em Go é usada para avaliar uma expressão e executar diferentes blocos de código com base nos valores resultantes dessa expressão. A sintaxe básica é a seguinte:

```go
switch expressao {
case valor1:
    // código a ser executado se expressao for igual a valor1
case valor2:
    // código a ser executado se expressao for igual a valor2
default:
    // código a ser executado se nenhum dos casos anteriores for verdadeiro
}
```

Por exemplo:

```go
diaDaSemana := 3
switch diaDaSemana {
case 1:
    fmt.Println("Segunda-feira")
case 2:
    fmt.Println("Terça-feira")
case 3:
    fmt.Println("Quarta-feira")
default:
    fmt.Println("Outro dia da semana")
}
```

**Expressões de Caso**

Em Go, os casos em um switch não precisam ser apenas valores constantes. Você pode usar expressões em cada caso, o que oferece grande flexibilidade. Por exemplo:

```go
diaDaSemana := "quarta"
switch diaDaSemana {
case "segunda", "terça", "quarta", "quinta", "sexta":
    fmt.Println("Dia útil")
case "sábado", "domingo":
    fmt.Println("Fim de semana")
default:
    fmt.Println("Dia inválido")
}
```

**Expressões Vazias**

Se a expressão de um switch for omitida, o switch é avaliado como true e o primeiro caso que resultar em true será executado. Isso permite criar estruturas de controle mais concisas e legíveis. Por exemplo:

```go
hora := 12
switch {
case hora < 12:
    fmt.Println("Bom dia!")
case hora >= 12 && hora < 18:
    fmt.Println("Boa tarde!")
default:
    fmt.Println("Boa noite!")
}
```

**Fallthrough**

Em Go, o comportamento padrão de um switch é sair após a execução de um caso. No entanto, você pode usar a palavra-chave fallthrough para forçar a execução do próximo caso, independentemente da expressão de correspondência. Por exemplo:

```go
switch num := 10; {
case num < 20:
    fmt.Println("Menor que 20")
    fallthrough
case num < 15:
    fmt.Println("Menor que 15")
}
```

Neste exemplo, "Menor que 20" é impresso e, em seguida, "Menor que 15" também é impresso devido ao fallthrough.
