package main

import (
	"fmt"
)

func diasDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	}
	return "Número Inválido"
}

func diasDaSemana2(numero int) string {

	retorno := ""

	switch numero {
	case 1:
		retorno = "Domingo"
		fallthrough
	case 2:
		retorno = "Segunda-Feira"
	case 3:
		retorno = "Terça-Feira"
	case 4:
		retorno = "Quarta-Feira"
	case 5:
		retorno = "Quinta-Feira"
	case 6:
		retorno = "Sexta-Feira"
	case 7:
		retorno = "Sabado"
	}
	return retorno
}

func main() {
	fmt.Println("Switch")

	fmt.Println(diasDaSemana(3))

	fmt.Println(diasDaSemana2(1))

}
