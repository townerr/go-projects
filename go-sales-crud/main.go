package main

import (
	"go-sales-crud/config"
	"go-sales-crud/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.Connect()
}

func main() {
	config.Connect()
	r := gin.New()

	routes.Routes(r)

	r.Run(":8080")
}
