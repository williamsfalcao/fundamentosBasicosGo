package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrmentando i", i)
		time.Sleep(time.Second)

	}
	fmt.Println(i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrmentando j", j)
		time.Sleep(time.Second)
	}

	//loop range
	nomes := [3]string{"Joao", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(letra)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
