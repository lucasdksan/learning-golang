package main

import (
	"command_line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start")

	aplication := app.Generate()
	erro := aplication.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}
