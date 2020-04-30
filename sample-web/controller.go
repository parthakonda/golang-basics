package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Compute(c *gin.Context) {
	min := 4000
	max := 8000
	waitTime := rand.Intn(max-min) + min

	fmt.Println(waitTime)
	for a := 0; a < waitTime; a++ {
		compute() // to put cpu intense stuff
	}
	c.JSON(http.StatusOK, "Ok")
}
