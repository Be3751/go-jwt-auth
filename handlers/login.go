package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/be3/go-jwt-auth/crypto"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	ID  string `json:"id"`
	PWD string `json:"pwd"`
}

type LoginRes struct {
	Token string `json:"token"`
}

// ユーザの確認とトークンの生成
func Login(ctx *gin.Context) {
	fmt.Println("/login")

	var loginReq LoginReq
	ctx.BindJSON(&loginReq) // リクエストボディのパラメータを割り当て

	// 認証できればJWTを返す
	if authenticate(loginReq) {
		fmt.Println("Authenticated to make a token.")
		token, err := crypto.GenerateToken(loginReq.ID, time.Now())
		if err != nil {
			fmt.Println("Couldn't generate a jwt.: ", err)
			ctx.Status(http.StatusInternalServerError)
		}
		loginRes := LoginRes{Token: token}
		ctx.JSON(http.StatusOK, loginRes)
	} else {
		fmt.Println("Invalid id or pwd.")
		ctx.Status(http.StatusUnauthorized)
	}
}

func authenticate(loginReq LoginReq) bool {
	mockRecord := LoginReq{ID: "123", PWD: "piyo"}
	if loginReq.ID == mockRecord.ID && loginReq.PWD == mockRecord.PWD {
		return true
	} else {
		return false
	}
}
