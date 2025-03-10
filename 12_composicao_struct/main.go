package main

import "fmt"

// declarando uma struct
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	//compondo uma struct com outra struct
	Endereco
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
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
	lucas.Cidade = "Porciuncula"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s\n", lucas.Nome, lucas.Idade, lucas.Ativo, lucas.Cidade)
}

/*
	- go nao é orientada a objetos
	- struct não é uma classe

	- compondo uma struc
		type Cliente struct {
		Nome  string
		Idade int
		Ativo bool
		Endereco
	}

	- criando uma propriedade do tipo endereço
		type Cliente struct {
			Nome   string
			Idade  int
			Ativo  bool
			Adress Endereco
	}

	- ex: lucas.Adress.Logradouro = "Rua 1"
*/
