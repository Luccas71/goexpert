package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d, cap=%d, %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d, cap=%d, %v\n", len(s[:2]), cap(s[:2]), s[:2])

	fmt.Printf("len=%d, cap=%d, %v\n", len(s[2:]), cap(s[2:]), s[2:])

	fmt.Printf("len=%d, cap=%d, %v\n", len(s[1:3]), cap(s[1:3]), s[1:3])

	// adiciona o item no fim do slice
	s = append(s, 12)
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)

}

/*
	SLICE:
	=> possiu ponteiro (apontando para um array), tamanho e capacidade
	=> tamanho dinâmico
	=> [:i] pega i valores à direita e ignorar o restante
	=> [i:] ignora i valores à esquerda e pega o restante à direita
	=> [i:j] pega os itens no intervalo determinado, incluindo valor de i
	=> sempre que o tamanho inicial do slice estoura, a capacidade INICIAL é DOBRADA
*/
