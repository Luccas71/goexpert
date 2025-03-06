package main

//importando um pacote externo
import (
	"fmt"
)

type ID int

var a ID = 12

func main() {
	// %T mostra o tipo da var
	fmt.Printf("O tipo de dado da var A é %T", a)
	//O tipo de dado da var A é main.ID => main.ID pq o tipo da var foi criado no pacote main
}
