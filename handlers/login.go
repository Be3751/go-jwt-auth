package handlers

import (
	"fmt"
	"net/http"

	"github.com/be3/go-jwt-auth/crypto"
	"github.com/gin-gonic/gin"
)

// ユーザの確認とトークンの生成
func Login(ctx *gin.Context) {
	fmt.Println("/login")

	var loginReq LoginReq
	ctx.BindJSON(&loginReq) // リクエストボディのパラメータを割り当て
	if true {
		fmt.Println("Authenticated to make a token.")
		token := crypto.CreateToken()
		ctx.String(http.StatusOK, token)
	} else {
		ctx.Status(http.StatusUnauthorized)
	}
}

type LoginReq struct {
	ID  string `json:"id"`
	PWD string `json:"pwd"`
}

type LoginRes struct {
	Token string `json:"token"`
}
