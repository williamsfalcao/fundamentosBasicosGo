package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Machado",
		"profissao": "Jogador de Futebol",
		"posicao":   "Atacante",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"Pedro": {
			"profissao": "Jogador de Futebol",
			"posicao":   "Atacante",
		},
		"Marina": {
			"profissao": "Medica",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "Pedro")
	fmt.Println(usuario2)
}
