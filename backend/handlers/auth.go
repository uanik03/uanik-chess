package handlers

import (
	"chess-backend/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context){

	user.Login("ani@gmail.com", "alo")

	c.JSON(http.StatusOK, gin.H{

		"message": "Hello, login",
	})

}