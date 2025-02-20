package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hellow from Go!"})
	})

	route.Run(":8081")
}
