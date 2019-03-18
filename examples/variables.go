package main

import "fmt"

func main() {
	var booleanVar bool
	var unsignedInt8Variable uint8
	var unsignedInt16Variable uint16
	var unsignedInt32variable uint32
	var unsignedInt64Variable uint64
	var unsignedIntVariable uint

	var i, j, k = 1, true, "test"
	dynamic := "Know"

	unsignedInt8Variable = 255

	fmt.Println(booleanVar)
	fmt.Println(unsignedInt8Variable)
	fmt.Println(unsignedInt16Variable)
	fmt.Println(unsignedInt32variable)
	fmt.Println(unsignedInt64Variable)
	fmt.Println(unsignedIntVariable)
	fmt.Println(i, j, k)
	fmt.Println(dynamic)
}
