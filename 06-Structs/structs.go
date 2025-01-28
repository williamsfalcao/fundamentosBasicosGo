package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {
	fmt.Println("Arquivo Strunct")

	endExemplo := endereco{"Rua dos Bobos", 0}

	var u usuario
	u.nome = "Williams"
	u.idade = 42
	fmt.Println(u.nome, u.idade)

	usuario2 := usuario{"Pedro", 8, endExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Marina"}
	fmt.Println(usuario3)

}
