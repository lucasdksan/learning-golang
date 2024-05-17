package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Faça um programa que abra um arquivo texto "input.txt", pré-existente. O programa então deve ler o arquivo linha por linha e apresentar a soma total do comprimento de todas as linhas. Verifique se é necessário fechar o arquivo antes do programa terminar. Adicione também suporte para eventuais situações de erro, como por exemplo não conseguir abrir o nome de arquivo especificado.

func main() {
	filename := "example.txt"

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var len_text_total uint64 = 0

	for scanner.Scan() {
		var text string = scanner.Text()
		len_text_total = len_text_total + uint64((len(text)))
	}

	fmt.Println("Total: ", len_text_total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
