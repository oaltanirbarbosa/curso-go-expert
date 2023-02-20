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
	fmt.Printf("O tipo de E é %T", f)
}
