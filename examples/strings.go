package main

import "fmt"
import "strings"

func main() {
	message := []string{"Welcome", "partha"}

	for i := 0; i < len(message); i++ {
		fmt.Printf("%x", message[i])
	}

	fmt.Printf("%s", message)
	fmt.Printf("%s", strings.Join(message, " "))
}
