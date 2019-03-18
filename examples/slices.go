/**
* Go slices are same as go arrays
* go arrays:
*	- store same type with fixed length
*   - you can't extend the size
*   - you can't get subslice of it
* go slices:
*   - you can extend the size of it using append
*   - you can get the subslice of the array
 */

package main

import "fmt"

func main() {
	var array = [5]int{}
	var slice = []int{1, 2, 3, 4, 5, 6, 7}
	var destination = make([]int, 8, 16)
	var nilslice []int
	array[0] = 1
	// slice[0] = 1 (this won't work, index out of range)
	slice = append(slice, 1)

	fmt.Print(array)
	fmt.Print(slice)

	x := slice[2:4]
	y := array[2:4]

	fmt.Print(x)
	fmt.Print(y)

	fmt.Println("Subslicing....")
	fmt.Println(array[:3])
	fmt.Println(slice[:3])
	/**
	* len() and cap()
	 */

	fmt.Print(len(array))
	fmt.Println(len(slice))

	fmt.Println(cap(array))
	fmt.Println(cap(slice))

	/**
	* using make
	 */
	dynamic := make([]int, 3, 5)

	fmt.Println("Dynamic array using slice")
	fmt.Print(len(dynamic), cap(dynamic), dynamic)

	/**
	* nil slice
	 */
	if dynamic == nil {
		fmt.Println("dynamic slice is nil")
	}
	if slice == nil {
		fmt.Println("slice slice is nil")
	}
	if nilslice == nil {
		fmt.Println("nilslice slice is nil")
	}

	/**
	* copy()
	 */
	copy(destination, slice)
	fmt.Println(destination)
}
