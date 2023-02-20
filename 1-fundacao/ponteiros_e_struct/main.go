package main

import "fmt"

// import "fmt"

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Wesley Willians"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
	// wesley := Cliente{
	// 	nome: "Wesley",
	// }
	// wesley.andou()
	// fmt.Printf("O valor da strinct com nome %v\n", wesley.nome)
}
