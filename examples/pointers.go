package main

import "fmt"

func main() {
	var a int
	var aptr *int
	var nilptr *int
	var dblptr **int

	var arrptr = []*int{}

	aptr = &a
	dblptr = &aptr

	fmt.Printf("address of a = %x", &a)
	fmt.Printf("value of aptr = %x", aptr)
	fmt.Printf("value of a = %d", a)
	fmt.Printf("value of a = %d", *aptr)
	fmt.Printf("value of nilptr %x", nilptr)
	fmt.Printf("address of aptr is = %x", &aptr)
	fmt.Printf("address of aptr dblptr = %x", dblptr)
	fmt.Printf("value of a = %d", **dblptr)

	if nilptr == nil {
		fmt.Printf("nilptr is not having values int it.")
	}

	arrptr = append(arrptr, &a)

	for _, item := range arrptr {
		fmt.Printf("item value - %x", item)
		fmt.Printf("item pointed value - %d", *item)
	}
}
