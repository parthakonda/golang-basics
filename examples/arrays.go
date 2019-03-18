package main

import "fmt"

func getSum(values []int) int {
	var sum int
	for _, val := range values {
		fmt.Printf("\n - %d", val)
		sum += val
	}
	return sum
}

func main() {
	/**
	Arrays can store data of same type
	*/
	var marks = []int{}
	var source = []int{1, 2, 3}
	destination := source

	marks = append(marks, 20)
	fmt.Print(marks)

	/**
	Initializing arrays
	*/
	var days = [7]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday"}
	for index, element := range days {
		fmt.Printf("%d day is - %s\n", index, element)
	}

	/**
	Accessing values at index
	*/
	fmt.Printf("\nvalue at location 2 is - %s", days[2])

	/**
	Two dimensional arrays
	*/
	var imageValues = [2][2]int{
		{0, 1},
		{2, 3}}
	for idx, i := range imageValues {
		for j, col := range imageValues[idx] {
			fmt.Printf("row - %d, col - %d, value: %d", i, j, col)
		}
	}
	fmt.Printf("sum of values is %d\n", getSum([]int{1, 2, 3}))

	fmt.Println("source = ", source)
	fmt.Println("destination = ", destination)
	// Modifying source
	source[0] = 5
	fmt.Println("source = ", source)
	fmt.Println("destination = ", destination)

	// Modifying destination
	destination[0] = 6
	fmt.Println("source = ", source)
	fmt.Println("destination = ", destination)

}
