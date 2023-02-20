package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

// func soma[T int | float64](m map[string]T) T {
func soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Wesley": 1000, "Joao": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 100.20, "Joao": 2000.3, "Maria": 300}
	m3 := map[string]MyNumber{"Wesley": 1000, "Joao": 2000, "Maria": 3000}
	println(soma(m))
	println(soma(m2))
	println(soma(m3))
	println(compara(10, 11))
}
