package main

import "fmt"

const a = "Hello Go"

type ID int

var (
	b bool = true
	c int
	d string
	e float64
	f ID = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 11
	meuArray[1] = 12
	meuArray[2] = 13
	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é igual a %d\n", i, v)
	}
}
