package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá mundo!")
	}()

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parametro")

	fmt.Println(retorno)
}
