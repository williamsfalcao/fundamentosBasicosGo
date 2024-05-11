package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 10000000
	fmt.Println(numero)

	var numero2 uint32 = 100000
	fmt.Println(numero2)

	//alias
	// INT32 = rune
	var numero3 rune = 100000
	fmt.Println(numero3)

	// INT8 = byte
	var numero4 byte = 100
	fmt.Println(numero4)

	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	numeroReal2 := 12345.67
	fmt.Println(numeroReal2)

	var str string = "Texto"
	fmt.Println(str)

	//Não existe char no go
	char := 'B'
	fmt.Println(char)

	var booleano bool
	fmt.Println(booleano)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}
