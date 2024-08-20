## Lista de Exercìcios

* 1 - Escreva um programa em Go que imprima "Olá, mundo!" na tela.

* 2 - Escreva um programa em Go que solicite ao usuário seu nome e, em seguida, imprima "Olá, [nome]!" na tela.

* 3 - Escreva uma função em Go que receba dois números como parâmetros e retorne a soma desses números.

* 4 - Escreva uma função em Go que receba um número inteiro como parâmetro e retorne verdadeiro se o número for par e falso caso contrário.

* 5 - Escreva uma função em Go que receba um número inteiro como parâmetro e retorne verdadeiro se o número for primo e falso caso contrário.

* 6 - Escreva uma função em Go que receba uma string como parâmetro e retorne o número de vogais nessa string.

* 7 - Escreva uma função em Go que receba uma lista de inteiros e retorne o maior número da lista.

* 8 - Escreva uma função em Go que receba uma lista de inteiros e retorne a média desses números.

* 9 - Escreva uma função em Go que receba uma lista de strings e retorne uma nova lista com todas as strings em letras maiúsculas.

* 10 - Escreva um programa em Go que leia um número inteiro do usuário e imprima todos os números primos menores ou iguais a esse número.

* 11 - Escreva uma função em Go que receba uma string e retorne verdadeiro se a string for um palíndromo (ou seja, pode ser lida da mesma forma de trás para frente) e falso caso contrário.

* 12 - Escreva uma função em Go que receba uma lista de inteiros e retorne uma nova lista contendo apenas os números pares da lista original.

* 13 - Escreva uma função em Go que calcule e retorne o fatorial de um número inteiro fornecido como parâmetro.

* 14 - Escreva uma função em Go que receba uma lista de strings e retorne a concatenação de todas as strings.

* 15 - Escreva uma função em Go que receba um número inteiro como parâmetro e retorne verdadeiro se o número for um quadrado perfeito e falso caso contrário.

* 16 - Escreva um programa em Go que imprima os primeiros N termos da sequência de Fibonacci, onde N é fornecido pelo usuário.

* 17 - Faça uma função/método que receba uma string como parâmetro e que retorne uma nova string, onde a sequência dos caracteres foi invertida. Dentro da parte principal (main), leia uma string digitada pelo usuário e passe para a função/método criada, imprimindo em seguida a string devolvida.

* 18 - Faça um programa que abra um arquivo texto "input.txt", pré-existente. O programa então deve ler o arquivo linha por linha e apresentar a soma total do comprimento de todas as linhas. Verifique se é necessário fechar o arquivo antes do programa terminar. Adicione também suporte para eventuais situações de erro, como por exemplo não conseguir abrir o nome de arquivo especificado.

* 19 - Escreva uma função em Go que receba uma lista de inteiros e retorne a soma de todos os números na lista.

* 20 - Escreva uma função em Go que receba uma string e retorne o número de palavras na string.

* 21 - Escreva uma função em Go que receba um número inteiro positivo como parâmetro e retorne uma lista de todos os divisores desse número.

* 22 - Escreva um programa em Go que solicite ao usuário uma temperatura em Celsius e imprima a temperatura equivalente em Fahrenheit.

* 23 - Escreva um programa em Go que faça o seguinte:

- 1 Crie um canal que transporta inteiros.
- 2 Inicie uma goroutine que envia números de 1 a 5 para esse canal.
- 2 No goroutine principal, receba e imprima esses números do canal.

* 24 - Escreva um programa em Go que faça o seguinte:

- 1 Crie dois canais que transportam strings.
- 2 Inicie duas goroutines que enviam mensagens diferentes para cada canal em intervalos diferentes.
- 3 Utilize um select para receber e imprimir as mensagens dos dois canais.

* 25 - Escreva um programa em Go que faça o seguinte:

- 1 Crie um canal com buffer de tamanho 3 que transporta inteiros.
- 2 Envie três valores para o canal sem utilizar goroutines.
- 3 Receba e imprima esses valores do canal.

* 26 - Produtor e Consumidor:

- 1 Crie um canal que transporta inteiros.
- 2 Inicie uma goroutine que atue como produtor, gerando números de 1 a 10 e enviando-os para o canal.
- 3 Inicie outra goroutine que atue como consumidor, recebendo os números do canal e imprimindo-os.
- 4 Use um canal adicional para sinalizar quando todos os números tiverem sido consumidos, e o goroutine principal deve esperar por esse sinal para terminar a execução.

* 27 - Goroutines Simples:

- 1 Crie duas goroutines.
- 2 Cada goroutine deve imprimir uma mensagem diferente.
- 3 Certifique-se de que o programa principal espere a conclusão das goroutines antes de terminar.

* 28 - Produtor-Consumidor com Buffer:

- 1 Crie um canal com buffer de tamanho 5.
- 2 Inicie uma goroutine produtora que envia 10 números inteiros para o canal.
- 3 Inicie uma goroutine consumidora que lê os números do canal e imprime-os.
- 4 Certifique-se de que o programa principal espere a conclusão das goroutines antes de terminar.

* 29 - Uso do select com Timeout:

- 1 Crie um canal que transporta inteiros.
- 2 Inicie uma goroutine que envia um número para o canal após um segundo.
- 3 Use select no goroutine principal para tentar receber do canal com um timeout de 500 milissegundos.
- 4 Imprima uma mensagem apropriada se a operação de recebimento no canal expirar.

* 30 - Paralelismo com WaitGroup

- 1 Crie 5 goroutines, cada uma imprimindo seu número identificador.
- 2 Utilize sync.WaitGroup para garantir que o programa principal espere todas as goroutines terminarem antes de encerrar.

* 31 - Mutex para Sincronização

- 1 Crie uma variável compartilhada (contador).
- 2 Inicie 10 goroutines que incrementam essa variável 100 vezes cada uma.
- 3 Use sync.Mutex para garantir que os incrementos no contador sejam realizados de forma segura.
- 4 Imprima o valor final do contador no goroutine principal.
