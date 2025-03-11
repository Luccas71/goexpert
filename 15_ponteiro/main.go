package main

func main() {

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
*/
