package main

import (
	"fmt"
)

func main() {
	//Aritimeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	//Atribuição
	var variavael1 string = "String"
	variavael2 := "String2"

	fmt.Println(variavael1, variavael2)

	//Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	fmt.Println("---------------")

	//Logicos
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true && true)

	//Unários
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	fmt.Println(numero)

	//Ternarios não existem em GOLang

}
