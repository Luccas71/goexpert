package main

import "fmt"

func main() {
	// array => estrutura de dados que possiu tamanho fixo e pode ser percorrido
	var arr [3]int
	arr[0] = 12
	arr[1] = 123
	arr[2] = 71

	//pegando o ultimo item do array
	fmt.Println(arr[len(arr)-1])

	//criando um for para percorrer o array
	for i, v := range arr {
		fmt.Printf("O indice é %d e o valor é %d\n", i, v)
	}
}


























