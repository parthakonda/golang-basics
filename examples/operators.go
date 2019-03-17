package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c int
	var _true, _false = true, false
	var ptr *int

	fmt.Printf("a = %d, b = %d", a, b)

	/**
	Aritchmetic Operators
	*/
	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - b = %d\n", a-b)
	fmt.Printf("a * b = %d\n", a*b)
	fmt.Printf("a / b = %d\n", a/b)
	fmt.Printf("a /% b = %d\n", b%a)

	/**
	Relational Operators
	*/
	fmt.Printf("a == b = %t\n", a == b)
	fmt.Printf("a != b = %t\n", a != b)
	fmt.Printf("a < b = %t\n", a < b)
	fmt.Printf("a > b = %t\n", a > b)
	fmt.Printf("a <= b = %t\n", a <= b)
	fmt.Printf("a >= b = %t\n", a >= b)

	/**
	Logical Operators
	*/
	fmt.Printf("a && b = %t\n", (_true && _false))
	fmt.Printf("a || b = %t\n", (_true || _false))
	fmt.Printf("!(a && b) = %t\n", !(_true && _false))

	/**
	Assignment Operators
	*/
	c += a
	fmt.Printf("c += a: %d\n", c)
	c -= a
	fmt.Printf("c -= a: %d\n", c)
	c *= b
	fmt.Printf("c *= b: %d\n", c)
	c /= a
	fmt.Printf("c /= a: %d\n", c)
	c %= 2
	fmt.Printf("c %= 2: %d\n", c)

	/**
	Address & Pointer type
	*/
	ptr = &a
	fmt.Printf("ptr = ", ptr)
	fmt.Printf("&ptr = ", &ptr)
	fmt.Printf("*ptr = ", *ptr)

}
