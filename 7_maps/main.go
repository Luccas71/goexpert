package main

import "fmt"

func main() {

	salarios := map[string]int{"Lucas": 1200, "Cadu": 1000}
	fmt.Println(salarios["Lucas"])

	//deletando uma chave
	delete(salarios, "Cadu")
	fmt.Println(salarios)

	//adicionando uma chave nova
	salarios["Joao"] = 2000
	fmt.Println(salarios)

	//iniciando maps vazios
	// sal := make(map[string]int)
	// sal2 := map[string]int{}

	//percorrendo um map
	for nome, salario := range salarios {
		fmt.Printf("O salario do %s é %d\n", nome, salario)
	}

	// usando _ (blank identifier) para ignorar uma variavel 
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
