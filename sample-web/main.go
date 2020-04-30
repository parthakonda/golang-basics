package main

import (
	"log"
)

func main() {
	router.GET("/", Compute)
	log.Fatal(router.Run(":8080"))
}
