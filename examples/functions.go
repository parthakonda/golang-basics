package main

import "fmt"

func displayName(name string, age int) {
	fmt.Printf("Your name is - %s\nyour age is - %d", name, age)
}

func add(a int, b int) int {
	return a + b
}

/**
Returning multiple values
*/
func swap(a int, b int) (int, int) {
	return b, a
}

/**
Function arguments `call by value`
*/
func swapNumbers(a int, b int) int {
	var temp int
	temp = a
	a = b
	b = temp
	return temp
}

/**
Function arguments `call by reference`
*/
func swapNumbersRef(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

/**
Closure functions
*/

func getNextNumber() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Numbers struct {
	a int
	b int
}

func (numbers Numbers) add() int {
	return numbers.a + numbers.b
}

func (numbers Numbers) sub() int {
	return numbers.a - numbers.b
}

func main() {
	var x, y = 1, 2
	var start, end = 1, 20
	displayName("Partha", 26)
	fmt.Printf("addition of 1 + 3 is = %d", add(1, 3))
	a, b := swap(1, 2)
	fmt.Printf("swap of 1, 2 is %d %d", a, b)

	// Before swapping values of x, y
	fmt.Printf("Before swapping...")
	fmt.Printf("x = %d\ny = %d", x, y)
	swapNumbers(x, y)
	fmt.Printf("Call by value: After swapping...")
	fmt.Printf("x = %d\ny = %d", x, y)
	swapNumbersRef(&x, &y)
	fmt.Printf("Call by Reference: After swapping...")
	fmt.Printf("x = %d\ny = %d", x, y)

	/**
	Function as value
	*/
	myFunction := func(start int, end int) int {
		return end - start
	}

	difference := myFunction(start, end)
	fmt.Printf("different between start - %d, end - %d is %d", start, end, difference)

	nextNumber := getNextNumber()

	fmt.Printf("Next number %d\n", nextNumber())
	fmt.Printf("Next number %d\n", nextNumber())
	fmt.Printf("Next number %d\n", nextNumber())

	nextNumber = getNextNumber()

	fmt.Printf("Next number %d\n", nextNumber())
	fmt.Printf("Next number %d\n", nextNumber())
	fmt.Printf("Next number %d\n", nextNumber())

	numbers := Numbers{20, 10}
	fmt.Printf("add = %d", numbers.add())
	fmt.Printf("subtract = %d", numbers.sub())
}
