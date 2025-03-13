package main

import "fmt"

var p *int

func main() {
	x := 10
	p = &x

	// fmt.Println(x, p)

	//acessando o valor atraves do ponteiro
	// fmt.Println(*p)

	//modificando um valor atraves do ponteiro
	*p = 20
	fmt.Println(x, *p)
}


/*
	- ponteiro -> o ponteiro guarda a localização de um valor na memória.

	- declaração de um ponteiro:
		var p *int
		p é um ponteiro para um valor do tipo int

	- Atribuição de um Endereço a um Ponteiro:
		x := 10
		p = &x
		p agora guarda o endereço de memória da variável x

	- Acessando o Valor Através do Ponteiro:
		fmt.Println(*p) // Imprime 10
		Aqui, *p significa "o valor que está no endereço guardado por p".

	- Modificando o Valor Através do Ponteiro:
		*p = 20
		fmt.Println(x) // Imprime 20
		Aqui, *p = 20 altera o valor de x para 20, porque p aponta para x.


	=> Resumo
		Ponteiros armazenam endereços de memória.

		Use & para obter o endereço de uma variável.

		Use * para acessar ou modificar o valor no endereço guardado pelo ponteiro.

		Ponteiros são úteis para modificar variáveis fora de funções, trabalhar com structs grandes e verificar valores nil.
*/
