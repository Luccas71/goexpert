package main

import "fmt"

type Cliente struct {
	Nome string
}

func (c Cliente) Andou() {
	fmt.Printf("O cliente %v andou\n", c.Nome)
}

type Conta struct {
	Saldo int
}

// alterando o valor real na conta
func (c *Conta) simular(valor int) int {
	c.Saldo += valor
	println(c.Saldo)
	return c.Saldo
}

// em todo lugar que a conta for alterada, será refletido no valor real dela
func NewConta() *Conta{
	return &Conta{Saldo: 0}
}
func main() {
	// lucas := Cliente{
	// 	Nome: "Lucas",
	// }
	// lucas.Andou()

	conta := Conta{
		Saldo: 120,
	}
	conta.simular(200)
	println(conta.Saldo)
}
