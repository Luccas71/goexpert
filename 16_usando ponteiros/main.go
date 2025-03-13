package main

// func vai receber dois ponteiros(ou seja, endereços na memoria)
func somar(a, b *int) int {
	//vai alterar o valor do ponteiro da variavel a
	*a = 40
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	//recebe o endereço na memoria das variaveis
	//quando for chamada irá alterar o valor na memoria da variavel no local de a
	somar(&minhaVar1, &minhaVar2)
	println(minhaVar1)
}

/*
	- Quando usar ponteiros?
		quando desejo alterar o valor na memoria e nao apenas fazer uma copia desse valor
*/
