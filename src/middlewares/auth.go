package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	jwt_service "github.com/rayyanhunerkar/todo-go/src/security/jwt"
	"github.com/spf13/viper"
)

func AuthJWTMiddleware(conf *viper.Viper) gin.HandlerFunc {
	jwtConf := jwt_service.InitJWTConf(conf)
	return func(ctx *gin.Context) {
		token, err := jwt_service.JWTService.ValidateToken(jwtConf, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "User is not authenticated")
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id := claims["id"]
			ctx.Set(
				"currentUser", id,
			)
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "User is not authenticated")
			return
		}
	}
}
