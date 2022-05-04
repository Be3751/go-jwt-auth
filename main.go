package main

import (
	"github.com/be3/go-jwt-auth/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)
	router.GET("/who-you-are", handlers.WhoYouAre)
	router.Run(":3000")
}
