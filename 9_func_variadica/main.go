package main

import "fmt"

func main() {
	fmt.Println(sum(12, 25, 358, 154, 568))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

/*
	função variádica pode receber varios argumentos desde que sejam do type permitido pela função
	(numeros ...int)
*/
