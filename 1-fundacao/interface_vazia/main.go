package main

import "fmt"

func main() {
	var minhaVar interface{} = "Wesley Willians"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("o valor de res é %v e o resultado de ok é %v\n", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v\n", res2)

	// var x interface{} = 10
	// var y interface{} = "Hello World"
	// showType(x)
	// showType(y)
}

// func showType(t interface{}) {
// 	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
// }
