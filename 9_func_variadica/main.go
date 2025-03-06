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

/*
=> Parâmetro
	É definido na declaração da função
	Indica os dados que a função espera receber
	Define o tipo e o número de entradas que uma função pode aceitar
	É utilizado quando está sendo criado um método
=> Argumento
	É o valor real que é passado para a função ou método quando ele é chamado
	É utilizado para executar métodos que esteja parametrizado
	Tem que ser do mesmo tipo que foi utilizado no parâmetro

*/
