package main

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
	a := "x"
	// a:= "fd"
	println(a, b, c, d, e, f)
}
