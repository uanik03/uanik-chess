package main

import (
	"chess-backend/config"
	"chess-backend/handlers"
	"chess-backend/internal/api"
	

	"github.com/gin-gonic/gin"

	"net/http"
)

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Middleware logic before request

		c.Next()

		// Middleware logic after request

	}

}

func main() {

	router := gin.Default()

	wss := api.NewWebSocketServer()
	config.Init()

	router.Use(Logger())

	router.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{

			"message": "Hello, Gin with Middleware!",
		})

	})

	router.GET("/ws", wss.HandleWebsocketConnection)

	auth := router.Group("/auth")
	{
		auth.GET("/login", handlers.LoginHandler)

	}
	user := router.Group("/user")
	{
		user.GET("/getUser", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{

				"message": "Hello, getuser",
			})

		})
	}

	router.Run(":8080")

}
