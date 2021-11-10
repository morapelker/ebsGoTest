package main

import (
	"fmt"
	"net/http"
)
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "d",
		})
	})

	r.GET("/mor", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"result": "ok"})
	})
	fmt.Println("server started")
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
