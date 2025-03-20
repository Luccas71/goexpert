package main

import "fmt"

func main() {
	var minhaVar interface{} = "Lucas"

	// vai tentar imprimir sem saber o tipo e nao vai dar certo
	println(minhaVar)

	//inferindo o tipo
	println(minhaVar.(string))

	res, ok := minhaVar.(int64)
	fmt.Printf("O valor de res é %d e o valor de ok é %v\n", res, ok)
}
