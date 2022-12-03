package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		m := "go-deployment-samples-gin-devops"
		c.JSON(http.StatusOK, gin.H{"message": m})
	})

	return r
}

func main() {
	r := setupRouter()
	// In production set GIN_MODE to release
	// or add this line:
	// gin.SetMode(gin.ReleaseMode)
	print("Server listening on 0.0.0.0:8080")
	r.Run(":8080")
}
