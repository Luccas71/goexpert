package main

// só é permitido passar MÉTODOS
// deve ser exatamente igual a assinatura do método
type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

type Empresa struct {
	Nome string
}

// qualquer struct que tiver o método Desativar() será uma Pessoa
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

// implementa a interface pessoa
func (e *Empresa) Desativar() {}

// implementa a interface pessoa
func (c *Cliente) Desativar() {
	c.Ativo = false
}

func main() {

}

/*
	- toda struct que possiu um método da interface, automaticamente implementa essa interface
*/
