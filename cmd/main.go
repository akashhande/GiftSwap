package main

import (
	"GittSwap/pkg/api"
	"GittSwap/pkg/schema"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	schema.Init()

	api.RegisterRoutes(r)

	// Start the server
	r.Run(":8080")
}
