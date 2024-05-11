package main

import "fmt"

func main() {

	fmt.Println("Estrutura de Controle")

	numero := 9

	if numero > 15 {
		fmt.Println("Número maior que 15")
	} else {
		fmt.Println("Número menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero < 10 {
		fmt.Println("Número menor que 10")
	}

}
