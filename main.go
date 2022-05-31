package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	tr := router.Group("/transacoes")
	tr.GET("/list", getAll)
	tr.POST("/new", newTransaction)

	router.Run()
}
