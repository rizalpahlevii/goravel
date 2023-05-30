package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/http/middleware"

	"goravel/app/http/controllers"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(http.StatusOK, http.Json{
			"Hello": "Goravel",
		})
	})

	userController := controllers.NewUserController()
	authController := controllers.NewAuthController()
	profileController := controllers.NewProfileController()
	facades.Route.Middleware(middleware.JwtMiddleware()).Get("/me", profileController.Me)
	facades.Route.Post("/login", authController.Login)
	facades.Route.Post("register", userController.Register)
	facades.Route.Get("/users/{id}", userController.Show)
}
