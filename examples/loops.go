package main

import "fmt"

func main() {
	var i = 0
	var number = [3]int{0, 1, 2}
	for a := 0; a < 10; a++ {
		fmt.Printf("a = %d\n", a)
	}

	for ; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	for i < 20 {
		i++
		fmt.Printf("i = %d\n", i)
	}

	for j, x := range number {
		fmt.Printf("x = %d at %d", x, j)
	}

	/**
	Nested loops
	*/
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}

	/**
	Break statement
	*/
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 3 {
				break
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}

	/**
	Continue statement
	*/
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 3 {
				continue
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}
