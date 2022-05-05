package handlers

import (
	"fmt"
	"net/http"

	"github.com/be3/go-jwt-auth/model"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		fmt.Println("Couldn't bind json.: ", err)
		ctx.Status(http.StatusInternalServerError)
	}
	if err := user.Create(); err != nil {
		fmt.Println("Couldn't create a user.: ", err)
		ctx.Status(http.StatusInternalServerError)
	}
	ctx.Status(http.StatusOK)
}
