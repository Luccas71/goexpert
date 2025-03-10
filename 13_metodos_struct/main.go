package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

// linkando um metodo a uma struct
func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {
	lucas := Cliente{
		Nome:  "Lucas",
		Idade: 30,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", lucas.Nome, lucas.Idade, lucas.Ativo)

	lucas.Desativar()
	fmt.Printf("Ativo: %t\n", lucas.Ativo)
}
