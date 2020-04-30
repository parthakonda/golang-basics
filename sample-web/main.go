package main

import (
	"log"
)

func main() {
	router.POST("/login", Login)
	log.Fatal(router.Run(":8080"))
}
