package main

import "fmt"

type Endereco struct {
	logradouro string
	numero     int
	cidade     string
	estado     string
}

type Cliente struct {
	nome     string
	idade    int
	ativo    bool
	endereco Endereco
}

func (c Cliente) Desativar() {
	c.ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.nome)
}

func main() {
	wesley := Cliente{
		nome:  "Wesley",
		idade: 30,
		ativo: true,
	}
	wesley.Desativar()
	// wesley.ativo = false
	wesley.endereco.cidade = "Salvador"
	fmt.Println(wesley.ativo)
	// fmt.Println(wesley.endereco.cidade)
	// fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.nome, wesley.idade, wesley.ativo)
}
