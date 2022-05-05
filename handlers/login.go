package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/be3/go-jwt-auth/crypto"
	"github.com/be3/go-jwt-auth/model"

	"github.com/gin-gonic/gin"
)

type RespToken struct {
	Token string `json:"token"`
}

// ユーザの確認とトークンの生成
func Login(ctx *gin.Context) {
	fmt.Println("/login")

	// リクエストボディのIDとPWDを割り当て
	var loginReq model.User
	ctx.BindJSON(&loginReq)

	// 認証できればJWTを返す
	if model.Authenticate(loginReq) {
		fmt.Println("Authenticated to make a token.")
		token, err := crypto.GenerateToken(loginReq.ID, time.Now())
		if err != nil {
			fmt.Println("Couldn't generate a jwt.: ", err)
			ctx.Status(http.StatusInternalServerError)
		}
		RespToken := RespToken{Token: token}
		ctx.JSON(http.StatusOK, RespToken)
	} else {
		fmt.Println("Invalid id or pwd.")
		ctx.Status(http.StatusUnauthorized)
	}
}
