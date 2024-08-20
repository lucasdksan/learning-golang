## Operadores

**Operadores Aritméticos**

Os operadores aritméticos em Go são usados para realizar operações matemáticas básicas. Os principais operadores aritméticos são:

```
+ : Adição
- : Subtração
* : Multiplicação
/ : Divisão
% : Módulo (resto da divisão)
```

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

```
== : Igual a
!= : Diferente de
< : Menor que
> : Maior que
<= : Menor ou igual a
>= : Maior ou igual a
```

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

```
& : E bit a bit (AND)
| : Ou bit a bit (OR)
^ : Ou exclusivo bit a bit (XOR)
<< : Deslocamento para a esquerda
>> : Deslocamento para a direita
```

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

```
+= : Adição e atribuição
-= : Subtração e atribuição
*= : Multiplicação e atribuição
/= : Divisão e atribuição
%= : Módulo e atribuição
<<= : Deslocamento para a esquerda e atribuição
>>= : Deslocamento para a direita e atribuição
&= : E bit a bit e atribuição
|= : Ou bit a bit e atribuição
^= : Ou exclusivo bit a bit e atribuição
```

Por exemplo:

```go
a := 10
a += 5   // a agora é 15
```

**Operador Ternário (Não Suportado em Go)**

É importante observar que Go não tem um operador ternário como algumas outras linguagens (por exemplo, ? : em C/C++). Em vez disso, geralmente se usa a estrutura if ou uma expressão condicional para alcançar o mesmo resultado.
