package main

import "fmt"

func main() {
	total := func() int {
		return sum(12, 43, 54)
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

/*
	closure => função anônima

*/
