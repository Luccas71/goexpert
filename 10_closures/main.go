package main

import "fmt"

func main() {
	nextInt := intSeq()

	fmt.Println("Chamada", nextInt())
	fmt.Println("Chamada", nextInt())
	fmt.Println("Chamada", nextInt())

	nextInt2 := intSeq()

	fmt.Println("Chamada", nextInt2())
	fmt.Println("Chamada", nextInt2())
	fmt.Println("Chamada", nextInt2())
}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/*
	closure => retorna uma função anônima
	=> armazena o estado
	=> o estado é exclusivo daquela chamada
	=> a cada nova atribuição, a função recomeça do seu estado inicial
*/
