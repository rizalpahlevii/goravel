package controllers

import "github.com/goravel/framework/contracts/http"

type ProfileController struct {
}

func NewProfileController() *ProfileController {
	return &ProfileController{}
}

func (r *ProfileController) Me(ctx http.Context) {
	ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Profile me",
		"success": true,
		"data":    nil,
	})
}
