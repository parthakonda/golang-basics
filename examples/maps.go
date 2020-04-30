package main

import "fmt"

func main() {
	dict := make(map[string]int);
	dict["key"] = 11
	dict["newkey"] = 12
	fmt.Println(dict)
	fmt.Println(dict["key"])
	delete(dict, "newkey")
	fmt.Println(dict)
}