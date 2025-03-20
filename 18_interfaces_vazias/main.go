package main

import "fmt"

func main() {
	var x interface{} = 12
	var y interface{} = true
	var z interface{} = "true"

	showType(x)
	showType(y)
	showType(z)
}

func showType(t interface{}) {
	fmt.Printf("O tipo do dado é %T\n", t)
}

/*
	aceita qualquer tipo de dado
	utilizar com sabedoria
	obs.: utilizar generics
*/
