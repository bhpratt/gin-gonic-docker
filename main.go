package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginCommand struct {
	Passcode string `json:"passcode"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello!",
		})
	})
	r.POST("/login", func(c *gin.Context) {
		var loginCmd LoginCommand
		c.BindJSON(&loginCmd)
		c.JSON(http.StatusOK, gin.H{"your passcode": loginCmd.Passcode})
	})
	r.Run(":8080") // listen and serve on :8080
}
