package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/paulohub/curso-go/matematica"
)

func main() {
	s := matematica.Soma(1, 2)
	fmt.Printf("resultado: %v\n", s)
	println(matematica.A)
	fmt.Println(uuid.New())
}
