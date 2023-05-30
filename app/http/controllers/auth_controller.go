package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/models"
	"goravel/app/service"
	"strconv"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (r *AuthController) Login(ctx http.Context) {
	var user models.User
	err := facades.Orm.Query().Where("email", ctx.Request().Input("email")).First(&user)
	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "User not found",
		})
		return
	}

	if user == (models.User{}) {
		ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "User not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	if helpers.VerifyPassword(user.Password, ctx.Request().Input("password")) == false {
		ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"message": "Password incorrect",
			"success": false,
			"data":    nil,
		})
		return
	}

	token, _ := service.GenerateToken(strconv.Itoa(user.Id))
	user.Token = token

	ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Login success",
		"success": true,
		"data":    user,
	})
}

func (r *AuthController) Logout(ctx http.Context) {
	ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Logout success",
		"success": true,
		"data":    nil,
	})
}
