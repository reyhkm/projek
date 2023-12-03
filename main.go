package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User credentials
const (
	ValidUsername = "reykal"
	ValidPassword = "alhikam123"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == ValidUsername && password == ValidPassword {
			// Login success, redirect to your website
			c.Redirect(http.StatusSeeOther, "https://asistenkelas.reyhzb.repl.co/")
		} else {
			// Login failed
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Invalid credentials",
			})
		}
	})

	router.Run(":8080")
}
