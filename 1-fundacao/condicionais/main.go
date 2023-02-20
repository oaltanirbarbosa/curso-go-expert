package main

func main() {
	a := 4
	b := 2
	c := 5
	if a > b && c > a {
		println("a > b && c > a")
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("d")
	}

	if a > b {
		println(a)
	} else {
		println(b)
	}
}
