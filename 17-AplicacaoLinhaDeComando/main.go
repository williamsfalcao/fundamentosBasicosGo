package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de Partida")

	aplicacao := app.Gerar()
	error := aplicacao.Run(os.Args)

	if error != nil {
		log.Fatal(error)
	}
}
