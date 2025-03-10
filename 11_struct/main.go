package main

import "fmt"

// declarando uma struct
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	//declarando um "objeto" do tipo cliente
	lucas := Cliente{
		Nome:  "Lucas",
		Idade: 30,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", lucas.Nome, lucas.Idade, lucas.Ativo)

	//mudando um valor
	lucas.Ativo = false
	lucas.Idade = 31

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", lucas.Nome, lucas.Idade, lucas.Ativo)
}

/*
	- go nao é orientada a objetos
	- struct não é uma classe
*/
