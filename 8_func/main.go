package main

import (
	"errors"
	"fmt"
)

func main() {
	//fazendo uma validação simples
	valor, err := sum(22, 34)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	if a+b > 50 {
		return 0, errors.New("O valor é maior que 50")
	}
	return a + b, nil
}

/*
	func sum(a, b int) int {
	return a + b
	}

	sum => nome da função
	a, b int => nome e tipo dos dados de entrada
	int => tipo de retorno da função
	return a + b => corpo da função

	uma função pode ter MAIS de um retorno (type, type)
*/
