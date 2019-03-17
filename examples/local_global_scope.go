package main

import "fmt"

var c int

func main() {
	var a, b int

	a = 10
	b = 20

	c = a + b
	fmt.Printf("a = %d\nb = %d\nc = %d", a, b, c)
}
