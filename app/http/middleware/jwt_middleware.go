package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/goravel/framework/contracts/http"
)

func JwtMiddleware() http.Middleware {
	return func(ctx http.Context) {
		tokenString := ctx.Request().Header("Authorization")
		if tokenString == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "Unauthorized",
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, nil
			}
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "Unauthorized",
			})
			return
		}

		ctx.Request().Next()

	}
}
