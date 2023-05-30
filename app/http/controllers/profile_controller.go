package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	"goravel/app/service"
)

type ProfileController struct {
}

func NewProfileController() *ProfileController {
	return &ProfileController{}
}

func (r *ProfileController) Me(ctx http.Context) {
	authorization := ctx.Request().Header("Authorization")
	payloads, _ := service.ExtractClaims(authorization)
	var user models.User
	err := facades.Orm.Query().Where("id", payloads["user_id"]).First(&user)
	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": err.Error(),
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

	ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Profile me",
		"success": true,
		"data":    user,
	})
}
