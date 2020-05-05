package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router.Run(":"+port)
}
