package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 1000, "Joã0": 2000, "Maria": 3000}
	// fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["Wes"] = 5000
	// fmt.Println(salarios)
	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}
	//black identifier
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
