package main

//const nao pode ter seu valor alterado
const a = "hello"

// o tipo de uma variavel declarada nao pode ser mudado
// var de escopo global

var (
	//declarando e atribuindo valor a uma var
	// go sempre infere um tipo para var

	b bool    = true
	c int     = 10
	d string  = "pri"
	e float64 = 1.3
)

func main() {

	// variavel de escopo local, só pode ser acessada dentro da função
	var a = 10
	println(a)

	// shorthand para atribuir var
	// : utilizado apenas quando declara a var
	b := 12
	println(b)
}
