package main

// import "fmt"

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	var1 := 10
	var2 := 20
	soma(&var1, &var2)
	println(var1)
	// a := 10ß
	// // &a endereco de memoria
	// // println(&a)
	// var ponteiro *int = &a
	// *ponteiro = 20
	// b := &a
	// *b = 30
	// println(a, b, *b)
}
