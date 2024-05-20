package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados.\n", u.nome)
}

func (u usuario) maiorIdade() {

	if u.idade >= 18 {
		fmt.Printf("Usuário %s é maior de idade.\n", u.nome)
	} else {
		fmt.Printf("Usuário %s é menor de idade.\n", u.nome)
	}

}

func (u *usuario) fazerAniversario() {
	u.idade++
	fmt.Println("Feliz Aniversário")
}

func main() {
	usuario := usuario{"Davi", 17}
	fmt.Println(usuario)
	usuario.salvar()
	usuario.maiorIdade()
	usuario.fazerAniversario()
	usuario.maiorIdade()
}
