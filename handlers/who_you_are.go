package handlers

import (
	"fmt"
	"net/http"

	"github.com/be3/go-jwt-auth/crypto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func WhoYouAre(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	fmt.Println(tokenString)
	// 検証済みトークンの取得
	token, err := crypto.VerifyToken(tokenString)
	if err != nil {
		fmt.Println("Invalid token.: ", err)
		c.Status(http.StatusUnauthorized)
	}
	claims := token.Claims.(jwt.MapClaims)
	message := fmt.Sprintf("You are %s", claims["user_id"])
	c.String(http.StatusOK, message)
}
